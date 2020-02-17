// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pkg/cli/deploy.go

// Package mocks is a generated GoMock package.
package mocks

import (
	archer "github.com/aws/amazon-ecs-cli-v2/internal/pkg/archer"
	deploy "github.com/aws/amazon-ecs-cli-v2/internal/pkg/deploy"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockenvironmentDeployer is a mock of environmentDeployer interface
type MockenvironmentDeployer struct {
	ctrl     *gomock.Controller
	recorder *MockenvironmentDeployerMockRecorder
}

// MockenvironmentDeployerMockRecorder is the mock recorder for MockenvironmentDeployer
type MockenvironmentDeployerMockRecorder struct {
	mock *MockenvironmentDeployer
}

// NewMockenvironmentDeployer creates a new mock instance
func NewMockenvironmentDeployer(ctrl *gomock.Controller) *MockenvironmentDeployer {
	mock := &MockenvironmentDeployer{ctrl: ctrl}
	mock.recorder = &MockenvironmentDeployerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockenvironmentDeployer) EXPECT() *MockenvironmentDeployerMockRecorder {
	return m.recorder
}

// DeployEnvironment mocks base method
func (m *MockenvironmentDeployer) DeployEnvironment(env *deploy.CreateEnvironmentInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeployEnvironment", env)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeployEnvironment indicates an expected call of DeployEnvironment
func (mr *MockenvironmentDeployerMockRecorder) DeployEnvironment(env interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeployEnvironment", reflect.TypeOf((*MockenvironmentDeployer)(nil).DeployEnvironment), env)
}

// StreamEnvironmentCreation mocks base method
func (m *MockenvironmentDeployer) StreamEnvironmentCreation(env *deploy.CreateEnvironmentInput) (<-chan []deploy.ResourceEvent, <-chan deploy.CreateEnvironmentResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StreamEnvironmentCreation", env)
	ret0, _ := ret[0].(<-chan []deploy.ResourceEvent)
	ret1, _ := ret[1].(<-chan deploy.CreateEnvironmentResponse)
	return ret0, ret1
}

// StreamEnvironmentCreation indicates an expected call of StreamEnvironmentCreation
func (mr *MockenvironmentDeployerMockRecorder) StreamEnvironmentCreation(env interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StreamEnvironmentCreation", reflect.TypeOf((*MockenvironmentDeployer)(nil).StreamEnvironmentCreation), env)
}

// DeleteEnvironment mocks base method
func (m *MockenvironmentDeployer) DeleteEnvironment(projName, envName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEnvironment", projName, envName)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEnvironment indicates an expected call of DeleteEnvironment
func (mr *MockenvironmentDeployerMockRecorder) DeleteEnvironment(projName, envName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEnvironment", reflect.TypeOf((*MockenvironmentDeployer)(nil).DeleteEnvironment), projName, envName)
}

// MockpipelineDeployer is a mock of pipelineDeployer interface
type MockpipelineDeployer struct {
	ctrl     *gomock.Controller
	recorder *MockpipelineDeployerMockRecorder
}

// MockpipelineDeployerMockRecorder is the mock recorder for MockpipelineDeployer
type MockpipelineDeployerMockRecorder struct {
	mock *MockpipelineDeployer
}

// NewMockpipelineDeployer creates a new mock instance
func NewMockpipelineDeployer(ctrl *gomock.Controller) *MockpipelineDeployer {
	mock := &MockpipelineDeployer{ctrl: ctrl}
	mock.recorder = &MockpipelineDeployerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockpipelineDeployer) EXPECT() *MockpipelineDeployerMockRecorder {
	return m.recorder
}

// CreatePipeline mocks base method
func (m *MockpipelineDeployer) CreatePipeline(env *deploy.CreatePipelineInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePipeline", env)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePipeline indicates an expected call of CreatePipeline
func (mr *MockpipelineDeployerMockRecorder) CreatePipeline(env interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePipeline", reflect.TypeOf((*MockpipelineDeployer)(nil).CreatePipeline), env)
}

// UpdatePipeline mocks base method
func (m *MockpipelineDeployer) UpdatePipeline(env *deploy.CreatePipelineInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePipeline", env)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePipeline indicates an expected call of UpdatePipeline
func (mr *MockpipelineDeployerMockRecorder) UpdatePipeline(env interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePipeline", reflect.TypeOf((*MockpipelineDeployer)(nil).UpdatePipeline), env)
}

// PipelineExists mocks base method
func (m *MockpipelineDeployer) PipelineExists(env *deploy.CreatePipelineInput) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PipelineExists", env)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PipelineExists indicates an expected call of PipelineExists
func (mr *MockpipelineDeployerMockRecorder) PipelineExists(env interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PipelineExists", reflect.TypeOf((*MockpipelineDeployer)(nil).PipelineExists), env)
}

// DeletePipeline mocks base method
func (m *MockpipelineDeployer) DeletePipeline(pipelineName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePipeline", pipelineName)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePipeline indicates an expected call of DeletePipeline
func (mr *MockpipelineDeployerMockRecorder) DeletePipeline(pipelineName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePipeline", reflect.TypeOf((*MockpipelineDeployer)(nil).DeletePipeline), pipelineName)
}

// AddPipelineResourcesToProject mocks base method
func (m *MockpipelineDeployer) AddPipelineResourcesToProject(project *archer.Project, region string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPipelineResourcesToProject", project, region)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPipelineResourcesToProject indicates an expected call of AddPipelineResourcesToProject
func (mr *MockpipelineDeployerMockRecorder) AddPipelineResourcesToProject(project, region interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPipelineResourcesToProject", reflect.TypeOf((*MockpipelineDeployer)(nil).AddPipelineResourcesToProject), project, region)
}

// GetProjectResourcesByRegion mocks base method
func (m *MockpipelineDeployer) GetProjectResourcesByRegion(project *archer.Project, region string) (*archer.ProjectRegionalResources, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjectResourcesByRegion", project, region)
	ret0, _ := ret[0].(*archer.ProjectRegionalResources)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProjectResourcesByRegion indicates an expected call of GetProjectResourcesByRegion
func (mr *MockpipelineDeployerMockRecorder) GetProjectResourcesByRegion(project, region interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectResourcesByRegion", reflect.TypeOf((*MockpipelineDeployer)(nil).GetProjectResourcesByRegion), project, region)
}

// GetRegionalProjectResources mocks base method
func (m *MockpipelineDeployer) GetRegionalProjectResources(project *archer.Project) ([]*archer.ProjectRegionalResources, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRegionalProjectResources", project)
	ret0, _ := ret[0].([]*archer.ProjectRegionalResources)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRegionalProjectResources indicates an expected call of GetRegionalProjectResources
func (mr *MockpipelineDeployerMockRecorder) GetRegionalProjectResources(project interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRegionalProjectResources", reflect.TypeOf((*MockpipelineDeployer)(nil).GetRegionalProjectResources), project)
}

// MockprojectDeployer is a mock of projectDeployer interface
type MockprojectDeployer struct {
	ctrl     *gomock.Controller
	recorder *MockprojectDeployerMockRecorder
}

// MockprojectDeployerMockRecorder is the mock recorder for MockprojectDeployer
type MockprojectDeployerMockRecorder struct {
	mock *MockprojectDeployer
}

// NewMockprojectDeployer creates a new mock instance
func NewMockprojectDeployer(ctrl *gomock.Controller) *MockprojectDeployer {
	mock := &MockprojectDeployer{ctrl: ctrl}
	mock.recorder = &MockprojectDeployerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockprojectDeployer) EXPECT() *MockprojectDeployerMockRecorder {
	return m.recorder
}

// DeployProject mocks base method
func (m *MockprojectDeployer) DeployProject(in *deploy.CreateProjectInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeployProject", in)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeployProject indicates an expected call of DeployProject
func (mr *MockprojectDeployerMockRecorder) DeployProject(in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeployProject", reflect.TypeOf((*MockprojectDeployer)(nil).DeployProject), in)
}

// AddAppToProject mocks base method
func (m *MockprojectDeployer) AddAppToProject(project *archer.Project, appName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAppToProject", project, appName)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAppToProject indicates an expected call of AddAppToProject
func (mr *MockprojectDeployerMockRecorder) AddAppToProject(project, appName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAppToProject", reflect.TypeOf((*MockprojectDeployer)(nil).AddAppToProject), project, appName)
}

// AddEnvToProject mocks base method
func (m *MockprojectDeployer) AddEnvToProject(project *archer.Project, env *archer.Environment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddEnvToProject", project, env)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEnvToProject indicates an expected call of AddEnvToProject
func (mr *MockprojectDeployerMockRecorder) AddEnvToProject(project, env interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEnvToProject", reflect.TypeOf((*MockprojectDeployer)(nil).AddEnvToProject), project, env)
}

// DelegateDNSPermissions mocks base method
func (m *MockprojectDeployer) DelegateDNSPermissions(project *archer.Project, accountID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DelegateDNSPermissions", project, accountID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DelegateDNSPermissions indicates an expected call of DelegateDNSPermissions
func (mr *MockprojectDeployerMockRecorder) DelegateDNSPermissions(project, accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelegateDNSPermissions", reflect.TypeOf((*MockprojectDeployer)(nil).DelegateDNSPermissions), project, accountID)
}

// DeleteProject mocks base method
func (m *MockprojectDeployer) DeleteProject(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProject", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProject indicates an expected call of DeleteProject
func (mr *MockprojectDeployerMockRecorder) DeleteProject(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProject", reflect.TypeOf((*MockprojectDeployer)(nil).DeleteProject), name)
}

// MockprojectResourcesGetter is a mock of projectResourcesGetter interface
type MockprojectResourcesGetter struct {
	ctrl     *gomock.Controller
	recorder *MockprojectResourcesGetterMockRecorder
}

// MockprojectResourcesGetterMockRecorder is the mock recorder for MockprojectResourcesGetter
type MockprojectResourcesGetterMockRecorder struct {
	mock *MockprojectResourcesGetter
}

// NewMockprojectResourcesGetter creates a new mock instance
func NewMockprojectResourcesGetter(ctrl *gomock.Controller) *MockprojectResourcesGetter {
	mock := &MockprojectResourcesGetter{ctrl: ctrl}
	mock.recorder = &MockprojectResourcesGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockprojectResourcesGetter) EXPECT() *MockprojectResourcesGetterMockRecorder {
	return m.recorder
}

// GetProjectResourcesByRegion mocks base method
func (m *MockprojectResourcesGetter) GetProjectResourcesByRegion(project *archer.Project, region string) (*archer.ProjectRegionalResources, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjectResourcesByRegion", project, region)
	ret0, _ := ret[0].(*archer.ProjectRegionalResources)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProjectResourcesByRegion indicates an expected call of GetProjectResourcesByRegion
func (mr *MockprojectResourcesGetterMockRecorder) GetProjectResourcesByRegion(project, region interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectResourcesByRegion", reflect.TypeOf((*MockprojectResourcesGetter)(nil).GetProjectResourcesByRegion), project, region)
}

// GetRegionalProjectResources mocks base method
func (m *MockprojectResourcesGetter) GetRegionalProjectResources(project *archer.Project) ([]*archer.ProjectRegionalResources, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRegionalProjectResources", project)
	ret0, _ := ret[0].([]*archer.ProjectRegionalResources)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRegionalProjectResources indicates an expected call of GetRegionalProjectResources
func (mr *MockprojectResourcesGetterMockRecorder) GetRegionalProjectResources(project interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRegionalProjectResources", reflect.TypeOf((*MockprojectResourcesGetter)(nil).GetRegionalProjectResources), project)
}

// Mockdeployer is a mock of deployer interface
type Mockdeployer struct {
	ctrl     *gomock.Controller
	recorder *MockdeployerMockRecorder
}

// MockdeployerMockRecorder is the mock recorder for Mockdeployer
type MockdeployerMockRecorder struct {
	mock *Mockdeployer
}

// NewMockdeployer creates a new mock instance
func NewMockdeployer(ctrl *gomock.Controller) *Mockdeployer {
	mock := &Mockdeployer{ctrl: ctrl}
	mock.recorder = &MockdeployerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Mockdeployer) EXPECT() *MockdeployerMockRecorder {
	return m.recorder
}

// DeployEnvironment mocks base method
func (m *Mockdeployer) DeployEnvironment(env *deploy.CreateEnvironmentInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeployEnvironment", env)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeployEnvironment indicates an expected call of DeployEnvironment
func (mr *MockdeployerMockRecorder) DeployEnvironment(env interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeployEnvironment", reflect.TypeOf((*Mockdeployer)(nil).DeployEnvironment), env)
}

// StreamEnvironmentCreation mocks base method
func (m *Mockdeployer) StreamEnvironmentCreation(env *deploy.CreateEnvironmentInput) (<-chan []deploy.ResourceEvent, <-chan deploy.CreateEnvironmentResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StreamEnvironmentCreation", env)
	ret0, _ := ret[0].(<-chan []deploy.ResourceEvent)
	ret1, _ := ret[1].(<-chan deploy.CreateEnvironmentResponse)
	return ret0, ret1
}

// StreamEnvironmentCreation indicates an expected call of StreamEnvironmentCreation
func (mr *MockdeployerMockRecorder) StreamEnvironmentCreation(env interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StreamEnvironmentCreation", reflect.TypeOf((*Mockdeployer)(nil).StreamEnvironmentCreation), env)
}

// DeleteEnvironment mocks base method
func (m *Mockdeployer) DeleteEnvironment(projName, envName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEnvironment", projName, envName)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEnvironment indicates an expected call of DeleteEnvironment
func (mr *MockdeployerMockRecorder) DeleteEnvironment(projName, envName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEnvironment", reflect.TypeOf((*Mockdeployer)(nil).DeleteEnvironment), projName, envName)
}

// DeployProject mocks base method
func (m *Mockdeployer) DeployProject(in *deploy.CreateProjectInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeployProject", in)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeployProject indicates an expected call of DeployProject
func (mr *MockdeployerMockRecorder) DeployProject(in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeployProject", reflect.TypeOf((*Mockdeployer)(nil).DeployProject), in)
}

// AddAppToProject mocks base method
func (m *Mockdeployer) AddAppToProject(project *archer.Project, appName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAppToProject", project, appName)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAppToProject indicates an expected call of AddAppToProject
func (mr *MockdeployerMockRecorder) AddAppToProject(project, appName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAppToProject", reflect.TypeOf((*Mockdeployer)(nil).AddAppToProject), project, appName)
}

// AddEnvToProject mocks base method
func (m *Mockdeployer) AddEnvToProject(project *archer.Project, env *archer.Environment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddEnvToProject", project, env)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEnvToProject indicates an expected call of AddEnvToProject
func (mr *MockdeployerMockRecorder) AddEnvToProject(project, env interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEnvToProject", reflect.TypeOf((*Mockdeployer)(nil).AddEnvToProject), project, env)
}

// DelegateDNSPermissions mocks base method
func (m *Mockdeployer) DelegateDNSPermissions(project *archer.Project, accountID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DelegateDNSPermissions", project, accountID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DelegateDNSPermissions indicates an expected call of DelegateDNSPermissions
func (mr *MockdeployerMockRecorder) DelegateDNSPermissions(project, accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelegateDNSPermissions", reflect.TypeOf((*Mockdeployer)(nil).DelegateDNSPermissions), project, accountID)
}

// DeleteProject mocks base method
func (m *Mockdeployer) DeleteProject(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProject", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProject indicates an expected call of DeleteProject
func (mr *MockdeployerMockRecorder) DeleteProject(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProject", reflect.TypeOf((*Mockdeployer)(nil).DeleteProject), name)
}

// CreatePipeline mocks base method
func (m *Mockdeployer) CreatePipeline(env *deploy.CreatePipelineInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePipeline", env)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePipeline indicates an expected call of CreatePipeline
func (mr *MockdeployerMockRecorder) CreatePipeline(env interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePipeline", reflect.TypeOf((*Mockdeployer)(nil).CreatePipeline), env)
}

// UpdatePipeline mocks base method
func (m *Mockdeployer) UpdatePipeline(env *deploy.CreatePipelineInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePipeline", env)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePipeline indicates an expected call of UpdatePipeline
func (mr *MockdeployerMockRecorder) UpdatePipeline(env interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePipeline", reflect.TypeOf((*Mockdeployer)(nil).UpdatePipeline), env)
}

// PipelineExists mocks base method
func (m *Mockdeployer) PipelineExists(env *deploy.CreatePipelineInput) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PipelineExists", env)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PipelineExists indicates an expected call of PipelineExists
func (mr *MockdeployerMockRecorder) PipelineExists(env interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PipelineExists", reflect.TypeOf((*Mockdeployer)(nil).PipelineExists), env)
}

// DeletePipeline mocks base method
func (m *Mockdeployer) DeletePipeline(pipelineName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePipeline", pipelineName)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePipeline indicates an expected call of DeletePipeline
func (mr *MockdeployerMockRecorder) DeletePipeline(pipelineName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePipeline", reflect.TypeOf((*Mockdeployer)(nil).DeletePipeline), pipelineName)
}

// AddPipelineResourcesToProject mocks base method
func (m *Mockdeployer) AddPipelineResourcesToProject(project *archer.Project, region string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPipelineResourcesToProject", project, region)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPipelineResourcesToProject indicates an expected call of AddPipelineResourcesToProject
func (mr *MockdeployerMockRecorder) AddPipelineResourcesToProject(project, region interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPipelineResourcesToProject", reflect.TypeOf((*Mockdeployer)(nil).AddPipelineResourcesToProject), project, region)
}

// GetProjectResourcesByRegion mocks base method
func (m *Mockdeployer) GetProjectResourcesByRegion(project *archer.Project, region string) (*archer.ProjectRegionalResources, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjectResourcesByRegion", project, region)
	ret0, _ := ret[0].(*archer.ProjectRegionalResources)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProjectResourcesByRegion indicates an expected call of GetProjectResourcesByRegion
func (mr *MockdeployerMockRecorder) GetProjectResourcesByRegion(project, region interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectResourcesByRegion", reflect.TypeOf((*Mockdeployer)(nil).GetProjectResourcesByRegion), project, region)
}

// GetRegionalProjectResources mocks base method
func (m *Mockdeployer) GetRegionalProjectResources(project *archer.Project) ([]*archer.ProjectRegionalResources, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRegionalProjectResources", project)
	ret0, _ := ret[0].([]*archer.ProjectRegionalResources)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRegionalProjectResources indicates an expected call of GetRegionalProjectResources
func (mr *MockdeployerMockRecorder) GetRegionalProjectResources(project interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRegionalProjectResources", reflect.TypeOf((*Mockdeployer)(nil).GetRegionalProjectResources), project)
}
