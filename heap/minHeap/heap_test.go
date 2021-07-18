package minHeap

import (
	"testing"
)

func TestInsert(t *testing.T) {
	h := NewMinHeap()
	h.Insert(7)
	h.Insert(3)
	h.Insert(1)
	if h.GetMin() != 1 {
		t.Error("expected min value to be 1 but got", h.GetMin())
	}
}

func TestRemoveMin(t *testing.T) {
	h := NewMinHeap()
	h.Insert(5)
	h.Insert(1)
	h.Insert(17)
	h.Insert(10)
	h.Insert(84)
	h.Insert(19)
	h.Insert(6)
	h.Insert(22)
	h.Insert(9)
	if h.RemoveMin() != 1 {
		t.Error("expected min value to be 1 but got", h.GetMin())
	}
	if h.GetMin() != 5 {
		t.Error("expected min value to be 5 but got", h.GetMin())
	}
}
