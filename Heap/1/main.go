// min-heap , build, insert, remove
package main

import "fmt"

type MinHeap struct {
	array []int
}

// Insert element to the heap
func (h *MinHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.heapifyUp(len(h.array) - 1)
}

// remove the minimum element
func (h *MinHeap) Remove() int {
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

// heapifyUp ensure that the heap is maintained while inserting
func (h *MinHeap) heapifyUp(index int) {
	for h.array[parent(index)] > h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// heapifydown ensure the heap is maintained while removing
func (h *MinHeap) heapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	smallest := index

	if l <= lastIndex && h.array[l] < h.array[smallest] {
		smallest = l
	}
	if r <= lastIndex && h.array[r] < h.array[smallest] {
		smallest = r
	}
	if smallest != index {
		h.swap(index, smallest)
		h.heapifyDown(smallest)
	}
}

// buildHeap builds a min-heap from an array
func (h *MinHeap) BuildHeap(arr []int) {
	h.array = arr
	for i := len(arr)/2 - 1; i >= 0; i-- {
		h.heapifyDown(i)
	}
}

// helper function
func (h *MinHeap) swap(i1, i2 int) {
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
	heap := &MinHeap{}
	heap.BuildHeap([]int{10, 20, 15, 30, 40})

	fmt.Println("Min-Heap array:", heap.array)

	heap.Insert(5)
	fmt.Println("After inserting 5:", heap.array)

	fmt.Println("Remove min:", heap.Remove())
	fmt.Println("After removing min:", heap.array)
}
