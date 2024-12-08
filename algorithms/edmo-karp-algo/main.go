package main

import (
	"fmt"

	"edmondkarp/graph"
)

func main() {
	// create a graph
	g := graph.NewGraph(5)

	// Add some edges
	g.AddEdge(0, 1, 10)
	g.AddEdge(1, 2, 5)
	g.AddEdge(0, 2, 15)

	// Add vertex data
	g.AddVertexData(0, "Source")
	g.AddVertexData(1, "Intermediate1")
	g.AddVertexData(2, "Sink")

	// Run EdmondsKarp(0, 2)
	maxFlow := g.EdmondsKarp
	fmt.Printf("Maximum Flow: %d\n", maxFlow)
}
