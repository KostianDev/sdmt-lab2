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

	t.Run("Insert elements", func(t *testing.T) {
		list := lists.NewList()
		err := list.Insert('a', 0)
		if err != nil || list.Length() != 1 {
			t.Errorf("Insert into empty list failed")
		}

		err = list.Insert('b', 0)
		if err != nil || list.Length() != 2 {
			t.Errorf("Insert at beginning failed")
		}

		err = list.Insert('c', 1)
		if err != nil || list.Length() != 3 {
			t.Errorf("Insert in middle failed")
		}

		err = list.Insert('d', 3)
		if err != nil || list.Length() != 4 {
			t.Errorf("Insert at end failed")
		}

		expected := []rune{'b', 'c', 'a', 'd'}
		for i, exp := range expected {
			val, err := list.Get(i)
			if err != nil || val != exp {
				t.Errorf("At index %d expected %c, got %c", i, exp, val)
			}
		}

		err = list.Insert('x', -1)
		if err == nil {
			t.Errorf("Expected error for negative index")
		}

		err = list.Insert('x', 5)
		if err == nil {
			t.Errorf("Expected error for index out of range")
		}
	})
}