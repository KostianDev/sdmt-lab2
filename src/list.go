package main

import "errors"

type List struct {
	elements []rune
}

func NewList() *List {
	return &List{
		elements: make([]rune, 0),
	}
}

func (l *List) Length() int {
	return len(l.elements)
}

func (l *List) Append(element rune) {
	l.elements = append(l.elements, element)
}

func (l *List) Insert(element rune, index int) error {
	if index < 0 || index > len(l.elements) {
		return errors.New("index out of bounds")
	}
	l.elements = append(l.elements[:index], append([]rune{element}, l.elements[index:]...)...)
	return nil
}

func (l *List) Delete(index int) (rune, error) {
	if index < 0 || index >= len(l.elements) {
		return 0, errors.New("index out of bounds")
	}
	element := l.elements[index]
	l.elements = append(l.elements[:index], l.elements[index+1:]...)
	return element, nil
}

func (l *List) DeleteAll(element rune) {
	for i := 0; i < len(l.elements); i++ {
		if l.elements[i] == element {
			l.elements = append(l.elements[:i], l.elements[i+1:]...)
			i--
		}
	}
}

func (l *List) Get(index int) (rune, error) {
	if index < 0 || index >= len(l.elements) {
		return 0, errors.New("index out of bounds")
	}
	return l.elements[index], nil
}

func (l *List) FindFirst(element rune) int {
	for i, e := range l.elements {
		if e == element {
			return i
		}
	}
	return -1
}

func (l *List) FindLast(element rune) int {
	for i := len(l.elements) - 1; i >= 0; i-- {
		if l.elements[i] == element {
			return i
		}
	}
	return -1
}

func (l *List) Clear() {
	l.elements = make([]rune, 0)
}