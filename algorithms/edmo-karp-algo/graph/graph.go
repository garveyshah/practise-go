package graph

import (
	"fmt"
	"math"
	"strings"
)

type Graph struct {
	adjMatrix  [][]int
	size       int
	vertexData []string
}

// NewGraph creates and returns a new Graph instance
func NewGraph(size int) *Graph {
	adjMatrix := make([][]int, size)
	for i := range adjMatrix {
		adjMatrix[i] = make([]int, size)
	}
	return &Graph{
		adjMatrix:  adjMatrix,
		size:       size,
		vertexData: make([]string, size),
	}
}

// AddEdge adds weighted adge between two vertices
func (g *Graph) AddEdge(u, v, capacity int) {
	if u >= 0 && u < g.size && v >= 0 && v < g.size {
		g.adjMatrix[u][v] = capacity
	}
}

// AddVertexData sets data for a specific vertex
func (g *Graph) AddVertexData(vertex int, data string) {
	if vertex >= 0 && vertex < g.size {
		g.vertexData[vertex] = data
	}
}

// BFS performs a breath-first search from source s to target t
// Returns true if target is reachable, false otherwise parent slice
// is used to track the path
func (g *Graph) BFS(s, t int, parent []int) bool {
	// Initialize visited and parent slices
	visited := make([]bool, g.size)
	for i := range parent {
		parent[i] = -1
	}

	// Create a queue for BFS
	queue := []int{s}
	visited[s] = true

	// BFS traversal
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		// Check all adjacent vertices
		for ind, val := range g.adjMatrix[u] {
			// if vertex is not visited and has a path
			if !visited[ind] && val > 0 {
				queue = append(queue, ind)
				visited[ind] = true
				parent[ind] = u
			}
		}
	}

	// Return whether target vertex was visited
	return visited[t]
}

// EdmondsKarp implements the Edmonds-Karp algorithm for max flow
func (g *Graph) EdmondsKarp(source, sink int) int {
	// Create a copy of the adjacency matrix to avoid modifying the original
	residualGraph := make([][]int, g.size)
	for i := range g.adjMatrix {
		residualGraph[i] = make([]int, g.size)
		copy(residualGraph[i], g.adjMatrix[i])
	}

	// Initialize variables
	maxFlow := 0
	parent := make([]int, g.size)

	// Run the algorithm while augmenting paths exist
	for {
		// Reset parent for each iteration
		for i := range parent {
			parent[i] = -1
		}

		// Find augmenting path using BFS
		if !g.bfsMaxFlow(source, sink, parent, residualGraph) {
			break
		}

		// Find minimum flow along the path
		pathFlow := math.MaxInt32
		v := sink
		for v != source {
			u := parent[v]
			pathFlow = min(pathFlow, residualGraph[u][v])
			v = parent[v]
		}

		// Update residual graph
		v = sink
		for v != source {
			u := parent[v]
			residualGraph[u][v] -= pathFlow
			residualGraph[v][u] += pathFlow
			v = parent[v]
		}

		// Reconstruct and print the path
		path := []int{}
		v = sink
		for v != source {
			path = append(path, v)
			v = parent[v]
		}
		path = append(path, source)

		// Reverse the path
		for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
			path[i], path[j] = path[j], path[i]
		}

		// Print path with vertex names
		pathNames := make([]string, len(path))
		for i, node := range path {
			pathNames[i] = g.vertexData[node]
		}
		fmt.Printf("Path: %s, Flow: %d\n", strings.Join(pathNames, " -> "), pathFlow)

		// Add to max flow
		maxFlow += pathFlow
	}

	return maxFlow
}

// bfsMaxFlow is a modified BFS for finding augmenting paths
func (g *Graph) bfsMaxFlow(source, sink int, parent []int, residualGraph [][]int) bool {
	visited := make([]bool, g.size)
	queue := []int{source}
	visited[source] = true
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		for v, capacity := range residualGraph[u] {
			if !visited[v] && capacity > 0 {
				queue = append(queue, v)
				parent[v] = u
				visited[v] = true

				if v == sink {
					return true
				}
			}
		}
	}
	return false
}

// helper function to find minimum of two intergers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
