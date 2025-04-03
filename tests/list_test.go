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
}