package main

type List struct {
	elements []rune
}

func NewList() *List {
	return &List{
		elements: make([]rune, 0),
	}
}