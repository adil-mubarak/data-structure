package main

import "fmt"

type MinHeap struct {
	array []int
}

func (h *MinHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.HeapIfUp(len(h.array) - 1)
}

func (h *MinHeap) Remove() int {
	if len(h.array) == 0 {
		fmt.Println("heap is empty")
		return -1
	}
	root := h.array[0]
	h.array[0] = h.array[len(h.array)-1]
	h.array = h.array[:len(h.array)-1]
	h.HeapIfDown(0)
	return root
}

func (h *MinHeap) HeapIfUp(index int) {
	for h.array[Parent(index)] > h.array[index] {
		h.Swap(Parent(index), index)
		index = Parent(index)
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
	h.array = arr
	for i := len(arr)/2 - 1; i >= 0; i-- {
		h.HeapIfDown(i)
	}
}

func (h *MinHeap) Swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func Parent(i int) int {
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
    fmt.Println("After Insert(5):", heap.array)

    fmt.Println("Remove min:", heap.Remove())
    fmt.Println("After removing min:", heap.array)

}
