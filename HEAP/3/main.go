// finding the kthlargest number
package main

import "fmt"

type MinHeap struct {
	array []int
	k     int
}

func (h *MinHeap) Insert(key int) {
	if len(h.array) < h.k {
		h.array = append(h.array, key)
		h.HeapIfUp(len(h.array) - 1)
	} else if key > h.array[0] {
		h.array[0] = key
		h.HeapIfDown(0)
	}
}

func (h *MinHeap) HeapIfUp(index int) {
	for index > 0 && h.array[parent(index)] > h.array[index] {
		h.Swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MinHeap) HeapIfDown(index int) {
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
		h.Swap(index, smallest)
		h.HeapIfDown(smallest)
	}
}

func (h *MinHeap) BuildHeap(arr []int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		h.HeapIfDown(i)
	}
}

func (h *MinHeap) Swap(i1, i2 int) {
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
	heap := &MinHeap{k: 3} 

	nums := []int{3, 1, 5, 12, 2, 11}
	for _, num := range nums {
		heap.Insert(num)
	}

	fmt.Println("3rd largest element:", heap.array[0])

}
