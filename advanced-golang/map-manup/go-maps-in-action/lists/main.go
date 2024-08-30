package main

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
}
