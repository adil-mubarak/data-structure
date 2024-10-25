// sorting another model
package main

import "fmt"

func HeapSort(arr []int) {
	n := len(arr)

	for i := n/2 - 1; i >= 0; i-- {
		HeapIfDown(arr, n, i)
	}

	for i := n - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		HeapIfDown(arr, i, 0)
	}
}

func HeapIfDown(arr []int, n, i int) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2

	if l < n && arr[l] > arr[largest] {
		largest = l
	}
	if r < n && arr[r] > arr[largest] {
		largest = r
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		HeapIfDown(arr, n, largest)
	}
}

func main() {
	arr := []int{12, 11, 13, 5, 6, 7}
	fmt.Println("Original array:", arr)

	HeapSort(arr)
	fmt.Println("Sorted array:", arr)

}
