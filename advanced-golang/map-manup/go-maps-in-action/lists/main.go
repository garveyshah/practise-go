package main

import "fmt"

type Node struct {
	Next  *Node
	Value interface{}
}

func main() {
	// Creating nodes
	node1 := &Node{Value: "Node 1"}
	node2 := &Node{Value: "Node 2"}
	node3 := &Node{Value: "Node 3"}

	// Linked nodes
	node1.Next = node2
	node2.Next = node3
	// Uncomment the  following line to create a cycle
	// node3.Next = node1

	// Starting point of the linked list
	var first *Node = node1

	// Map to track visited nodes
	visited := make(map[*Node]bool)

	// Traverse the list and detect cycles
	for n := first; n != nil; n = n.Next {
		if visited[n] {
			fmt.Println("Cycle detected")
			break
		}
		visited[n] = true
		fmt.Println(n.Value)
	}
}
