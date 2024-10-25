package main

import "fmt"

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Insert(value int) {
	h.array = append(h.array, value)
	h.HeapIfyUp(len(h.array) - 1)
}

func (h *MaxHeap) Remove() int {
	if h.array == nil {
		return 0
	}
	root := h.array[0]
	h.array[0] = h.array[len(h.array)-1]
	h.array = h.array[:len(h.array)-1]
	h.HeapIfDown(0)
	return root

}

func (h *MaxHeap) HeapIfyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.Swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MaxHeap) HeapIfDown(index int) {
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
		h.Swap(largest, index)
		h.HeapIfDown(largest)
	}
}

func (h *MaxHeap) BuildHeap(arr []int) {
	h.array = arr
	for i := len(arr)/2 - 1; i >= 0; i-- {
		h.HeapIfDown(i)
	}
}

func (h *MaxHeap) Swap(i1, i2 int) {
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
	heap.BuildHeap([]int{10, 2, 45, 6, 8})
	fmt.Println("heap: ", heap.array)

	heap.Insert(1)
	fmt.Println("after inserting 1: ", heap.array)

	fmt.Println("Removing the max value: ", heap.Remove())
	fmt.Println("after removing the max value: ", heap.array)
}
