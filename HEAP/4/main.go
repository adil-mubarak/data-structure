// heap sorting using max-heap
package main

import (
	"fmt"
)

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) HeapSort(arr []int) {
	h.array = arr
	n := len(arr)
	h.BuildHeap(arr)
	for i := n - 1; i > 0; i-- {
		h.Swap(0, i)
		h.HeapIfDownRange(0,i-1)
	}
}

func(h *MaxHeap)HeapIfDownRange(index , endIndex int){
	l,r := left(index),right(index)
	largest := index

	if l <= endIndex && h.array[l] > h.array[largest]{
		largest = l
	}
	if r <= endIndex && h.array[r] > h.array[largest]{
		largest = r
	}
	if largest != index{
		h.Swap(index,largest)
		h.HeapIfDownRange(largest,endIndex)
	}
}

func (h *MaxHeap) BuildHeap(arr []int) {
	h.array = arr
	for i := len(arr)/2 - 1; i >= 0; i-- {
		h.HeapIfDown(i)
	}
}

func (h *MaxHeap) HeapIfDown(index int) {
	largest := index
	l, r := left(index), right(index)
	lastIndex := len(h.array) - 1

	if l <= lastIndex && h.array[l] > h.array[largest] {
		largest = l
	}
	if r <= lastIndex && h.array[r] > h.array[largest] {
		largest = r
	}
	if largest != index {
		h.Swap(index, largest)
		h.HeapIfDown(largest)
	}
}

func (h *MaxHeap) Swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func main() {
	heap := &MaxHeap{}
	arr := []int{12, 11, 13, 5, 6, 7}
	fmt.Println("Original array:", arr)

	heap.HeapSort(arr)
	fmt.Println("Sorted array:", arr)

}
