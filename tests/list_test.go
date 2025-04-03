package tests

import (
	"testing"
	"github.com/KostianDev/sdmt-lab2/src"
)

func TestList(t *testing.T) {
	t.Run("Empty list", func(t *testing.T) {
		list := lists.NewList()
		if list.Length() != 0 {
			t.Errorf("Expected length 0, got %d", list.Length())
		}
	})

	t.Run("Append elements", func(t *testing.T) {
		list := lists.NewList()
		list.Append('a')
		list.Append('b')
		list.Append('c')

		if list.Length() != 3 {
			t.Errorf("Expected length 3, got %d", list.Length())
		}

		val, err := list.Get(1)
		if err != nil || val != 'b' {
			t.Errorf("Expected 'b' at index 1, got %c (err: %v)", val, err)
		}
	})
}