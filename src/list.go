package lists

type Node struct {
	value rune
	next  *Node
	previous *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	length int
}

func NewList() *LinkedList {
	return &LinkedList{nil, nil, 0}
}