package lists

import "errors"

type Node struct {
	value    rune
	next     *Node
	previous *Node
}

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func NewList() *LinkedList {
	return &LinkedList{nil, nil, 0}
}

func (l *LinkedList) Length() int {
	return l.length
}

func (l *LinkedList) Append(element rune) {
	node := &Node{value: element}
	if l.tail != nil {
		l.tail.next = node
		node.previous = l.tail
	} else {
		l.head = node
	}
	l.tail = node
	l.length++
}

func (l *LinkedList) Insert(element rune, index int) error {
	if index < 0 || index > l.length {
		return errors.New("index out of range")
	}
	node := &Node{value: element}
	if index == 0 {
		node.next = l.head
		if l.head != nil {
			l.head.previous = node
		}
		l.head = node
		if l.tail == nil {
			l.tail = node
		}
	} else if index == l.length {
		l.Append(element)
		return nil
	} else {
		current := l.head
		for i := 0; i < index-1; i++ {
			current = current.next
		}
		node.next = current.next
		node.previous = current
		current.next = node
		if node.next != nil {
			node.next.previous = node
		}
	}
	l.length++
	return nil
}

func (l *LinkedList) Delete(index int) (rune, error) {
	if index < 0 || index >= l.length {
		return 0, errors.New("index out of range")
	}
	var deletedValue rune
	if index == 0 {
		deletedValue = l.head.value
		l.head = l.head.next
		if l.head != nil {
			l.head.previous = nil
		} else {
			l.tail = nil
		}
	} else if index == l.length-1 {
		deletedValue = l.tail.value
		l.tail = l.tail.previous
		if l.tail != nil {
			l.tail.next = nil
		} else {
			l.head = nil
		}
	} else {
		current := l.head
		for i := 0; i < index; i++ {
			current = current.next
		}
		deletedValue = current.value
		current.previous.next = current.next
		current.next.previous = current.previous
	}
	l.length--
	return deletedValue, nil
}

func (l *LinkedList) DeleteAll(element rune) {
	for current := l.head; current != nil; {
		if current.value == element {
			if current.previous != nil {
				current.previous.next = current.next
			} else {
				l.head = current.next
			}
			if current.next != nil {
				current.next.previous = current.previous
			} else {
				l.tail = current.previous
			}
			l.length--
		}
		current = current.next
	}
}

func (l *LinkedList) Get(index int) (rune, error) {
	if index < 0 || index >= l.length {
		return 0, errors.New("index out of range")
	}
	current := l.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current.value, nil
}

func (l *LinkedList) FindFirst(element rune) int {
	current := l.head
	for i := 0; current != nil; i++ {
		if current.value == element {
			return i
		}
		current = current.next
	}
	return -1
}

func (l *LinkedList) FindLast(element rune) int {
	current := l.tail
	for i := l.length - 1; current != nil; i-- {
		if current.value == element {
			return i
		}
		current = current.previous
	}
	return -1
}

func (l *LinkedList) Clear() {
	l.head = nil
	l.tail = nil
	l.length = 0
}