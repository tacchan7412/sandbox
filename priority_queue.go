package main

import (
	"fmt"
	"math"
)

type heap struct {
	data []int
	last int
}

func swap(h *heap, i int, j int) {
	tmp := h.data[i]
	h.data[i] = h.data[j]
	h.data[j] = tmp
	return
}

func (h *heap) insert(newData int) {
	h.last++
	h.data = append(h.data, newData)

	i := h.last
	for i > 0 {
		if h.data[(i-1)/2] > h.data[i] {
			swap(h, (i-1)/2, i)
			i = (i - 1) / 2
		} else {
			return
		}
	}
	return
}

func (h *heap) deletemin() (data int) {
	data = h.data[0]
	h.data[0] = h.data[h.last]
	h.last--
	h.data = h.data[:h.last+1]

	i := 0
	for i < h.last/2 {
		if h.data[i] > h.data[i*2+1] {
			if h.data[i*2+1] > h.data[i*2+2] {
				swap(h, i, i*2+2)
				i = i*2 + 2
			} else {
				swap(h, i, i*2+1)
				i = i*2 + 1
			}
		} else if h.data[i] > h.data[i*2+2] {
			swap(h, i, i*2+2)
		} else {
			return
		}
	}
	return
}

func (h *heap) printTree() {
	fmt.Print("\n")
	rowNum := 1
	for i, d := range h.data {
		fmt.Printf("%d ", d)
		if int(math.Pow(2, float64(rowNum))-2) == i && i != len(h.data)-1 {
			fmt.Print("\n")
			rowNum++
		}
	}
	fmt.Print("\n")
}

func main() {
	h := &heap{
		last: -1,
		data: []int{},
	}
	h.insert(31)
	h.printTree()
	h.insert(21)
	h.printTree()
	h.insert(48)
	h.printTree()
	h.insert(9)
	h.printTree()
	data := h.deletemin()
	fmt.Printf("\n%d is deleted.", data)
	h.printTree()
	h.insert(26)
	h.printTree()
	h.insert(13)
	h.printTree()
	data = h.deletemin()
	fmt.Printf("\n%d is deleted.", data)
	h.printTree()
}
