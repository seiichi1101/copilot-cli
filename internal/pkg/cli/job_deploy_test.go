// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cli

import (
	"errors"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/aws/copilot-cli/internal/pkg/docker/dockerengine"

	"github.com/aws/copilot-cli/internal/pkg/cli/mocks"
	"github.com/aws/copilot-cli/internal/pkg/config"
	"github.com/aws/copilot-cli/internal/pkg/manifest"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

type deployJobMocks struct {
	mockWs                 *mocks.MockwsJobDirReader
	mockimageBuilderPusher *mocks.MockimageBuilderPusher
	mockInterpolator       *mocks.Mockinterpolator
}

func TestJobDeployOpts_Validate(t *testing.T) {
	testCases := map[string]struct {
		inAppName string
		inEnvName string
		inJobName string

		mockWs    func(m *mocks.MockwsJobDirReader)
		mockStore func(m *mocks.Mockstore)

		wantedError error
	}{
		"no existing applications": {
			mockWs:    func(m *mocks.MockwsJobDirReader) {},
			mockStore: func(m *mocks.Mockstore) {},

			wantedError: errNoAppInWorkspace,
		},
		"with workspace error": {
			inAppName: "phonetool",
			inJobName: "resizer",
			mockWs: func(m *mocks.MockwsJobDirReader) {
				m.EXPECT().ListJobs().Return(nil, errors.New("some error"))
			},
			mockStore: func(m *mocks.Mockstore) {},

			wantedError: errors.New("list jobs in the workspace: some error"),
		},
		"with job not in workspace": {
			inAppName: "phonetool",
			inJobName: "resizer",
			mockWs: func(m *mocks.MockwsJobDirReader) {
				m.EXPECT().ListJobs().Return([]string{}, nil)
			},
			mockStore: func(m *mocks.Mockstore) {},

			wantedError: errors.New("job resizer not found in the workspace"),
		},
		"with unknown environment": {
			inAppName: "phonetool",
			inEnvName: "test",
			mockWs:    func(m *mocks.MockwsJobDirReader) {},
			mockStore: func(m *mocks.Mockstore) {
				m.EXPECT().GetEnvironment("phonetool", "test").
					Return(nil, errors.New("unknown env"))
			},

			wantedError: errors.New("get environment test configuration: unknown env"),
		},
		"successful validation": {
			inAppName: "phonetool",
			inJobName: "resizer",
			inEnvName: "test",
			mockWs: func(m *mocks.MockwsJobDirReader) {
				m.EXPECT().ListJobs().Return([]string{"resizer"}, nil)
			},
			mockStore: func(m *mocks.Mockstore) {
				m.EXPECT().GetEnvironment("phonetool", "test").
					Return(&config.Environment{Name: "test"}, nil)
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// GIVEN
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockWs := mocks.NewMockwsJobDirReader(ctrl)
			mockStore := mocks.NewMockstore(ctrl)
			tc.mockWs(mockWs)
			tc.mockStore(mockStore)
			opts := deployJobOpts{
				deployWkldVars: deployWkldVars{
					appName: tc.inAppName,
					name:    tc.inJobName,
					envName: tc.inEnvName,
				},
				ws:    mockWs,
				store: mockStore,
			}

			// WHEN
			err := opts.Validate()

			// THEN
			if tc.wantedError != nil {
				require.EqualError(t, err, tc.wantedError.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestJobDeployOpts_Ask(t *testing.T) {
	testCases := map[string]struct {
		inAppName  string
		inEnvName  string
		inJobName  string
		inImageTag string

		wantedCalls func(m *mocks.MockwsSelector)

		wantedJobName  string
		wantedEnvName  string
		wantedImageTag string
		wantedError    error
	}{
		"prompts for environment name and job names": {
			inAppName:  "phonetool",
			inImageTag: "latest",
			wantedCalls: func(m *mocks.MockwsSelector) {
				m.EXPECT().Job("Select a job from your workspace", "").Return("resizer", nil)
				m.EXPECT().Environment("Select an environment", "", "phonetool").Return("prod-iad", nil)
			},

			wantedJobName:  "resizer",
			wantedEnvName:  "prod-iad",
			wantedImageTag: "latest",
		},
		"don't call selector if flags are provided": {
			inAppName:  "phonetool",
			inEnvName:  "prod-iad",
			inJobName:  "resizer",
			inImageTag: "latest",
			wantedCalls: func(m *mocks.MockwsSelector) {
				m.EXPECT().Job(gomock.Any(), gomock.Any()).Times(0)
				m.EXPECT().Environment(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
			},

			wantedJobName:  "resizer",
			wantedEnvName:  "prod-iad",
			wantedImageTag: "latest",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// GIVEN
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockSel := mocks.NewMockwsSelector(ctrl)

			tc.wantedCalls(mockSel)
			opts := deployJobOpts{
				deployWkldVars: deployWkldVars{
					appName:  tc.inAppName,
					name:     tc.inJobName,
					envName:  tc.inEnvName,
					imageTag: tc.inImageTag,
				},
				sel: mockSel,
			}

			// WHEN
			err := opts.Ask()

			// THEN
			if tc.wantedError == nil {
				require.NoError(t, err)
				require.Equal(t, tc.wantedJobName, opts.name)
				require.Equal(t, tc.wantedEnvName, opts.envName)
				require.Equal(t, tc.wantedImageTag, opts.imageTag)
			} else {
				require.EqualError(t, err, tc.wantedError.Error())
			}
		})
	}
}

func TestJobDeployOpts_configureContainerImage(t *testing.T) {
	mockError := errors.New("mockError")
	mockManifest := []byte(`name: mailer
type: 'Scheduled Job'
image:
  build:
    dockerfile: path/to/Dockerfile
    context: path
on:
  schedule: "@daily"`)
	mockMftNoBuild := []byte(`name: mailer
type: 'Scheduled Job'
image:
  location: foo/bar
on:
  schedule: "@daily"`)
	mockMftBuildString := []byte(`name: mailer
type: 'Scheduled Job'
image:
  build: path/to/Dockerfile
on:
  schedule: "@daily"`)
	mockMftNoContext := []byte(`name: mailer
type: 'Scheduled Job'
image:
  build:
    dockerfile: path/to/Dockerfile
on:
  schedule: "@daily"`)

	tests := map[string]struct {
		inputSvc   string
		setupMocks func(mocks deployJobMocks)

		wantErr      error
		wantedDigest string
	}{
		"should return error if ws ReadFile returns error": {
			inputSvc: "mailer",
			setupMocks: func(m deployJobMocks) {
				gomock.InOrder(
					m.mockWs.EXPECT().ReadWorkloadManifest("mailer").Return(nil, mockError),
				)
			},
			wantErr: fmt.Errorf("read job %s manifest: %w", "mailer", mockError),
		},
		"should return error if interpolation fail": {
			inputSvc: "mailer",
			setupMocks: func(m deployJobMocks) {
				gomock.InOrder(
					m.mockWs.EXPECT().ReadWorkloadManifest(gomock.Any()).Return(mockManifest, nil),
					m.mockInterpolator.EXPECT().Interpolate(string(mockManifest)).Return("", mockError),
				)
			},
			wantErr: fmt.Errorf("interpolate environment variables for mailer manifest: %w", mockError),
		},
		"should return error if workspace methods fail": {
			inputSvc: "mailer",
			setupMocks: func(m deployJobMocks) {
				gomock.InOrder(
					m.mockWs.EXPECT().ReadWorkloadManifest(gomock.Any()).Return(mockManifest, nil),
					m.mockInterpolator.EXPECT().Interpolate(string(mockManifest)).Return(string(mockManifest), nil),
					m.mockWs.EXPECT().CopilotDirPath().Return("", mockError),
				)
			},
			wantErr: fmt.Errorf("get copilot directory: %w", mockError),
		},
		"success without building and pushing": {
			inputSvc: "mailer",
			setupMocks: func(m deployJobMocks) {
				gomock.InOrder(
					m.mockWs.EXPECT().ReadWorkloadManifest("mailer").Return(mockMftNoBuild, nil),
					m.mockInterpolator.EXPECT().Interpolate(string(mockMftNoBuild)).Return(string(mockMftNoBuild), nil),
					m.mockWs.EXPECT().CopilotDirPath().Times(0),
					m.mockimageBuilderPusher.EXPECT().BuildAndPush(gomock.Any(), gomock.Any()).Times(0),
				)
			},
		},
		"should return error if fail to build and push": {
			inputSvc: "mailer",
			setupMocks: func(m deployJobMocks) {
				gomock.InOrder(
					m.mockWs.EXPECT().ReadWorkloadManifest("mailer").Return(mockManifest, nil),
					m.mockInterpolator.EXPECT().Interpolate(string(mockManifest)).Return(string(mockManifest), nil),
					m.mockWs.EXPECT().CopilotDirPath().Return("/ws/root/copilot", nil),
					m.mockimageBuilderPusher.EXPECT().BuildAndPush(gomock.Any(), gomock.Any()).Return("", mockError),
				)
			},
			wantErr: fmt.Errorf("build and push image: mockError"),
		},
		"success": {
			inputSvc: "mailer",
			setupMocks: func(m deployJobMocks) {
				gomock.InOrder(
					m.mockWs.EXPECT().ReadWorkloadManifest("mailer").Return(mockManifest, nil),
					m.mockInterpolator.EXPECT().Interpolate(string(mockManifest)).Return(string(mockManifest), nil),
					m.mockWs.EXPECT().CopilotDirPath().Return("/ws/root/copilot", nil),
					m.mockimageBuilderPusher.EXPECT().BuildAndPush(gomock.Any(), &dockerengine.BuildArguments{
						Dockerfile: filepath.Join("/ws", "root", "path", "to", "Dockerfile"),
						Context:    filepath.Join("/ws", "root", "path"),
					}).Return("sha256:741d3e95eefa2c3b594f970a938ed6e497b50b3541a5fdc28af3ad8959e76b49", nil),
				)
			},
			wantedDigest: "sha256:741d3e95eefa2c3b594f970a938ed6e497b50b3541a5fdc28af3ad8959e76b49",
		},
		"using simple buildstring (backwards compatible)": {
			inputSvc: "mailer",
			setupMocks: func(m deployJobMocks) {
				gomock.InOrder(
					m.mockWs.EXPECT().ReadWorkloadManifest("mailer").Return(mockMftBuildString, nil),
					m.mockInterpolator.EXPECT().Interpolate(string(mockMftBuildString)).Return(string(mockMftBuildString), nil),
					m.mockWs.EXPECT().CopilotDirPath().Return("/ws/root/copilot", nil),
					m.mockimageBuilderPusher.EXPECT().BuildAndPush(gomock.Any(), &dockerengine.BuildArguments{
						Dockerfile: filepath.Join("/ws", "root", "path", "to", "Dockerfile"),
						Context:    filepath.Join("/ws", "root", "path", "to"),
					}).Return("sha256:741d3e95eefa2c3b594f970a938ed6e497b50b3541a5fdc28af3ad8959e76b49", nil),
				)
			},
			wantedDigest: "sha256:741d3e95eefa2c3b594f970a938ed6e497b50b3541a5fdc28af3ad8959e76b49",
		},
		"without context field in overrides": {
			inputSvc: "mailer",
			setupMocks: func(m deployJobMocks) {
				gomock.InOrder(
					m.mockWs.EXPECT().ReadWorkloadManifest("mailer").Return(mockMftNoContext, nil),
					m.mockInterpolator.EXPECT().Interpolate(string(mockMftNoContext)).Return(string(mockMftNoContext), nil),
					m.mockWs.EXPECT().CopilotDirPath().Return("/ws/root/copilot", nil),
					m.mockimageBuilderPusher.EXPECT().BuildAndPush(gomock.Any(), &dockerengine.BuildArguments{
						Dockerfile: filepath.Join("/ws", "root", "path", "to", "Dockerfile"),
						Context:    filepath.Join("/ws", "root", "path", "to"),
					}).Return("sha256:741d3e95eefa2c3b594f970a938ed6e497b50b3541a5fdc28af3ad8959e76b49", nil),
				)
			},
			wantedDigest: "sha256:741d3e95eefa2c3b594f970a938ed6e497b50b3541a5fdc28af3ad8959e76b49",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockWorkspace := mocks.NewMockwsJobDirReader(ctrl)
			mockimageBuilderPusher := mocks.NewMockimageBuilderPusher(ctrl)
			mockInterpolator := mocks.NewMockinterpolator(ctrl)
			mocks := deployJobMocks{
				mockWs:                 mockWorkspace,
				mockimageBuilderPusher: mockimageBuilderPusher,
				mockInterpolator:       mockInterpolator,
			}
			test.setupMocks(mocks)
			opts := deployJobOpts{
				deployWkldVars: deployWkldVars{
					name: test.inputSvc,
				},
				unmarshal:          manifest.UnmarshalWorkload,
				imageBuilderPusher: mockimageBuilderPusher,
				ws:                 mockWorkspace,
				newInterpolator: func(app, env string) interpolator {
					return mockInterpolator
				},
			}

			gotErr := opts.configureContainerImage()

			if test.wantErr != nil {
				require.EqualError(t, gotErr, test.wantErr.Error())
			} else {
				require.NoError(t, gotErr)
				require.Equal(t, test.wantedDigest, opts.imageDigest)
			}
		})
	}
}
