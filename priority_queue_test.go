package main

import "testing"

func TestSwap(t *testing.T) {
	h := &heap{
		data: []int{1, 2},
		last: 1,
	}
	swap(h, 0, 1)
	if h.data[0] != 2 {
		t.Errorf("failed to swap")
	}
	if h.data[1] != 1 {
		t.Errorf("failed to swap")
	}
}

func TestInsert(t *testing.T) {
	h := &heap{
		data: []int{1, 2},
		last: 1,
	}
	h.insert(3)
	if h.data[2] != 3 {
		t.Errorf("failed to insert")
	}
	if h.last != 2 {
		t.Errorf("failed to increment heap.last")
	}
	h = &heap{
		data: []int{5, 4, 3},
		last: 2,
	}
	h.insert(2)
	if h.data[0] != 2 {
		t.Errorf("failed to insert")
	}
	if h.data[1] != 5 {
		t.Errorf("failed to insert")
	}
	if h.data[3] != 4 {
		t.Errorf("failed to insert")
	}
}

func TestDeletemin(t *testing.T) {
	h := &heap{
		data: []int{13, 21, 48, 31, 26},
		last: 4,
	}
	data := h.deletemin()
	if data != 13 {
		t.Errorf("failed to deletemin")
	}
	if h.last != 3 {
		t.Errorf("failed to decrement heap.last")
	}
	if h.data[0] != 21 {
		t.Errorf("failed to deletemin")
	}
	if h.data[1] != 26 {
		t.Errorf("failed to deletemin")
	}
}
