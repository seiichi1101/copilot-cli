// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package graph provides functionality for directed graphs.
package graph

// nodeStatus denotes the visiting status of a node when running DFS in a graph.
type nodeStatus int

const (
	unvisited nodeStatus = iota + 1
	visiting
	visited
)

// Graph represents a directed graph.
type Graph struct {
	nodes map[string]neighbors
}

// Edge represents one edge of a directed graph.
type Edge struct {
	From string
	To   string
}

type neighbors map[string]bool

// New initiates a new Graph.
func New() *Graph {
	return &Graph{
		nodes: make(map[string]neighbors),
	}
}

// Add adds a connection between two Nodes.
func (g *Graph) Add(edge Edge) {
	fromNode, toNode := edge.From, edge.To
	// Add origin node if doesn't exist.
	if _, ok := g.nodes[fromNode]; !ok {
		g.nodes[fromNode] = make(neighbors)
	}
	// Add edge.
	g.nodes[fromNode][toNode] = true
}

type findCycleTempVars struct {
	status     map[string]nodeStatus
	nodeParent map[string]string
	cycleStart string
	cycleEnd   string
}

// IsAcyclic checks if the graph is acyclic. If not, return the first detected cycle.
func (g *Graph) IsAcyclic() ([]string, bool) {
	var cycle []string
	status := make(map[string]nodeStatus)
	for node := range g.nodes {
		status[node] = unvisited
	}
	temp := findCycleTempVars{
		status:     status,
		nodeParent: make(map[string]string),
	}
	// We will run a series of DFS in the graph. Initially all vertices are marked unvisited.
	// From each unvisited node, start the DFS, mark it visiting while entering and mark it visited on exit.
	// If DFS moves to a visiting node, then we have found a cycle. The cycle itself can be reconstructed using parent map.
	// See https://cp-algorithms.com/graph/finding-cycle.html
	for node := range g.nodes {
		if status[node] == unvisited && g.hasCycles(&temp, node) {
			for n := temp.cycleStart; n != temp.cycleEnd; n = temp.nodeParent[n] {
				cycle = append(cycle, n)
			}
			cycle = append(cycle, temp.cycleEnd)
			return cycle, false
		}
	}
	return nil, true
}

func (g *Graph) hasCycles(temp *findCycleTempVars, currNode string) bool {
	temp.status[currNode] = visiting
	for node := range g.nodes[currNode] {
		if temp.status[node] == unvisited {
			temp.nodeParent[node] = currNode
			if g.hasCycles(temp, node) {
				return true
			}
		} else if temp.status[node] == visiting {
			temp.cycleStart = currNode
			temp.cycleEnd = node
			return true
		}
	}
	temp.status[currNode] = visited
	return false
}
