package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListMerge(l1 *List, l2 *List) {
	for l2.Head != nil {
		ListPushBack(l1, l2.Head.Data)
		l2.Head = l2.Head.Next
	}
}

func ListPushBack(l *List, n interface{}) {
	node := &NodeL{Data: n}
	if l.Head == nil {
		l.Head = node
	} else {
		l.Tail.Next = node
	}
	l.Tail = node
}
