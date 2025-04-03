package main

import (
	"fmt"
)

func main() {
	list := NewList()
	demoList(list)
}

func demoList(list ListInterface) {
	list.Append('a')
	list.Append('b')
	list.Append('c')
	fmt.Println("After Append('a', 'b', 'c'):", listToString(list))
	
	list.Insert('x', 1)
	fmt.Println("After Insert('x', 1):", listToString(list))
	
	val, _ := list.Get(2)
	fmt.Println("Get(2):", string(val))
	
	fmt.Println("Length:", list.Length())
	
	fmt.Println("FindFirst('b'):", list.FindFirst('b'))
	
	list.Append('b')
	fmt.Println("After Append('b'):", listToString(list))
	fmt.Println("FindLast('b'):", list.FindLast('b'))
	
	deleted, _ := list.Delete(0)
	fmt.Println("Delete(0):", string(deleted), "->", listToString(list))
	
	list.DeleteAll('b')
	fmt.Println("After DeleteAll('b'):", listToString(list))
	
	clone := list.Clone()
	clone.Append('y')
	fmt.Println("Original:", listToString(list))
	fmt.Println("Clone with added 'y':", listToString(clone))
	
	list.Reverse()
	fmt.Println("After Reverse:", listToString(list))
	
	other := NewList()
	other.Append('1')
	other.Append('2')
	list.Extend(other)
	fmt.Println("After Extend with [1, 2]:", listToString(list))
	fmt.Println("Second list after operation:", listToString(other))
	
	list.Clear()
	fmt.Println("After Clear:", listToString(list))
}

func listToString(list ListInterface) string {
	result := "["
	for i := 0; i < list.Length(); i++ {
		val, _ := list.Get(i)
		if i > 0 {
			result += ", "
		}
		result += string(val)
	}
	result += "]"
	return result
}
