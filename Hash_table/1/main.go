package main

import "fmt"

func main() {
	//create hash map
	hashMap := make(map[string]int)

	//inserting values
	hashMap["apple"] = 100
	hashMap["banana"] = 50
	hashMap["orange"] = 80

	//retriveng values
	fmt.Println("apple price:", hashMap["apple"])
	fmt.Println("banana price:", hashMap["banana"])
	fmt.Println("orange price", hashMap["orange"])

	//update values
	hashMap["orange"] = 70
	fmt.Println("update orange price:", hashMap["orange"])

	//delete values
	delete(hashMap, "banana")
	fmt.Println("after delete banana price:", hashMap["banana"])

}
