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

	t.Run("Delete elements", func(t *testing.T) {
		list := lists.NewList()
		list.Append('a')
		list.Append('b')
		list.Append('c')
		list.Append('d')

		val, err := list.Delete(1)
		if err != nil || val != 'b' || list.Length() != 3 {
			t.Errorf("Delete from middle failed")
		}

		val, err = list.Delete(0)
		if err != nil || val != 'a' || list.Length() != 2 {
			t.Errorf("Delete from beginning failed")
		}

		val, err = list.Delete(1)
		if err != nil || val != 'd' || list.Length() != 1 {
			t.Errorf("Delete from end failed")
		}

		val, err = list.Delete(0)
		if err != nil || val != 'c' || list.Length() != 0 {
			t.Errorf("Delete last element failed")
		}

		_, err = list.Delete(-1)
		if err == nil {
			t.Errorf("Expected error for negative index")
		}

		_, err = list.Delete(0)
		if err == nil {
			t.Errorf("Expected error for empty list")
		}
	})

	t.Run("Delete all elements", func(t *testing.T) {
		list := lists.NewList()
		list.Append('a')
		list.Append('b')
		list.Append('a')
		list.Append('c')
		list.Append('a')
		list.Append('d')

		list.DeleteAll('a')
		if list.Length() != 3 {
			t.Errorf("Expected length 3 after DeleteAll, got %d", list.Length())
		}

		for i := 0; i < list.Length(); i++ {
			val, _ := list.Get(i)
			if val == 'a' {
				t.Errorf("Found 'a' at index %d after DeleteAll", i)
			}
		}

		initialLength := list.Length()
		list.DeleteAll('x')
		if list.Length() != initialLength {
			t.Errorf("Length changed after DeleteAll non-existing element")
		}
	})

	t.Run("Get elements", func(t *testing.T) {
		list := lists.NewList()
		list.Append('a')
		list.Append('b')
		list.Append('c')

		val, err := list.Get(0)
		if err != nil || val != 'a' {
			t.Errorf("Get(0) failed")
		}

		val, err = list.Get(1)
		if err != nil || val != 'b' {
			t.Errorf("Get(1) failed")
		}

		val, err = list.Get(2)
		if err != nil || val != 'c' {
			t.Errorf("Get(2) failed")
		}

		_, err = list.Get(-1)
		if err == nil {
			t.Errorf("Expected error for negative index")
		}

		_, err = list.Get(3)
		if err == nil {
			t.Errorf("Expected error for index out of range")
		}
	})

	t.Run("Find elements", func(t *testing.T) {
		list := lists.NewList()
		list.Append('a')
		list.Append('b')
		list.Append('a')
		list.Append('c')
		list.Append('a')
		list.Append('d')

		if list.FindFirst('a') != 0 {
			t.Errorf("FindFirst('a') failed")
		}

		if list.FindFirst('b') != 1 {
			t.Errorf("FindFirst('b') failed")
		}

		if list.FindLast('a') != 4 {
			t.Errorf("FindLast('a') failed")
		}

		if list.FindLast('d') != 5 {
			t.Errorf("FindLast('d') failed")
		}

		if list.FindFirst('x') != -1 {
			t.Errorf("FindFirst('x') should return -1")
		}

		if list.FindLast('x') != -1 {
			t.Errorf("FindLast('x') should return -1")
		}
	})

	t.Run("Clone list", func(t *testing.T) {
		list := lists.NewList()
		list.Append('a')
		list.Append('b')
		list.Append('c')

		clone := list.Clone()
		if clone.Length() != list.Length() {
			t.Errorf("Clone length mismatch")
		}

		list.Append('d')
		if clone.Length() == list.Length() {
			t.Errorf("Clone is not independent")
		}

		for i := 0; i < clone.Length(); i++ {
			origVal, _ := list.Get(i)
			cloneVal, _ := clone.Get(i)
			if origVal != cloneVal {
				t.Errorf("Value mismatch at index %d", i)
			}
		}
	})

	t.Run("Reverse list", func(t *testing.T) {
		list := lists.NewList()
		list.Append('a')
		list.Append('b')
		list.Append('c')
		list.Append('d')

		list.Reverse()

		expected := []rune{'d', 'c', 'b', 'a'}
		for i, exp := range expected {
			val, err := list.Get(i)
			if err != nil || val != exp {
				t.Errorf("After reverse, at index %d expected %c, got %c", i, exp, val)
			}
		}

		emptyList := lists.NewList()
		emptyList.Reverse()
		if emptyList.Length() != 0 {
			t.Errorf("Reversing empty list should do nothing")
		}

		singleList := lists.NewList()
		singleList.Append('x')
		singleList.Reverse()
		val, _ := singleList.Get(0)
		if val != 'x' || singleList.Length() != 1 {
			t.Errorf("Reversing single-element list should do nothing")
		}
	})

	t.Run("Clear list", func(t *testing.T) {
		list := lists.NewList()
		list.Append('a')
		list.Append('b')
		list.Append('c')

		list.Clear()
		if list.Length() != 0 {
			t.Errorf("Clear failed, length should be 0")
		}

		_, err := list.Get(0)
		if err == nil {
			t.Errorf("Expected error after Get on cleared list")
		}
	})
}