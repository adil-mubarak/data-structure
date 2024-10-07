package main

import "fmt"

func main() {
	var arr = [7]int{6, 5, 4, 3, 9, 8, 0}
	var t int = 17

	result(arr[:], t)
	var arr1 = [3]int{}
	beforeinsertion(arr1[:])
	afterinsertion(arr1[:])

	var arr2 = [5]int{1,2,3,4,5}
	beforedeletion(arr2[:])
	afterdeletion(arr2[0:3])

	fmt.Println(arr[1:4])
}

func result(arr []int, t int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == t {
				re := []int{arr[i], arr[j]}
				fmt.Println(re)
				return
			}
		}
	}
	fmt.Println("No pairs found")
}

/// ARRAY INSERTION

//array before insertion

func beforeinsertion(arr1 []int) {
	fmt.Println("array before insertion:")
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("[%d] = %d \n", i, arr1[i])
	}
}

// array insertion

func afterinsertion(arr1 []int) {
	fmt.Println("array after insertion:")
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i + 1
		fmt.Printf("[%d] = %d \n", i, arr1[i])
	}
}

///ARRAY DELETION

//array before deletion

func beforedeletion(arr2 []int){
	fmt.Println("array before deletion:")
	for i := 0; i < len(arr2); i++ {
		fmt.Printf("[%d]=%d \n",i,arr2[i])
		
	}
}

// array after deletion

func afterdeletion(arr2 []int){
	fmt.Println("array after deletion:")
	for i := 0; i < len(arr2); i++ {
		fmt.Printf("[%d] = %d \n",i,arr2[i])
	}
}