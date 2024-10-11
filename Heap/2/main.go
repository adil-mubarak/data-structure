// max heap
package main

import (
	"fmt"
)

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.heapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) remove() int {
	if len(h.array) == 0 {
		fmt.Println("Heap is empty")
		return -1
	}
	root := h.array[0]
	h.array[0] = h.array[len(h.array)-1]
	h.array = h.array[:len(h.array)-1]
	h.heapifyDown(0)
	return root
}

func (h *MaxHeap) heapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MaxHeap) heapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	largest := index

	if l <= lastIndex && h.array[l] > h.array[largest] {
		largest = l
	}
	if r <= lastIndex && h.array[r] > h.array[largest] {
		largest = r
	}
	if largest != index {
		h.swap(index, largest)
		h.heapifyDown(largest)
	}
}

func (h *MaxHeap) BuildHeap(arr []int) {
	h.array = arr
	for i := len(arr)/2 - 1; i >= 0; i-- {
		h.heapifyDown(i)
	}
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func parent(i int) int {
	return (i - 1) / 2
}
func left(i int) int {
	return 2*i + 1
}
func right(i int) int {
	return 2*i + 2
}

func main() {
	heap := &MaxHeap{}
	heap.BuildHeap([]int{10, 20, 15, 30, 40})

	fmt.Println("Max-heap array:", heap.array)

	heap.Insert(50)
	fmt.Println("After insert 50:", heap.array)

	fmt.Println("remove max:", heap.remove())
	fmt.Println("after remove max", heap.array)
}
