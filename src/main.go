package main

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
