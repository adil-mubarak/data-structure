package main

import "fmt"

func QuickSort(arr []int, low, high int) {
	if low < high {
		pivotIndex := Partition(arr, low, high)

		QuickSort(arr, low, pivotIndex-1)
		QuickSort(arr, pivotIndex+1, high)
	}
}

func Partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			Swap(arr,i, j)
		}
	}
	Swap(arr,i+1, high)
	return i + 1
}

func Swap(arr []int,i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	fmt.Println("Unsorted array:", arr)

	QuickSort(arr, 0, len(arr)-1)
	fmt.Println("Sorted array:", arr)
}
