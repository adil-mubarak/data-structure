package main

import "fmt"

func main() {
	hashMap := make(map[string]int)

	hashMap["apple"] = 50
	hashMap["banana"] = 40
	hashMap["orange"] = 60

	fmt.Println("Apple price:", hashMap["apple"])
	fmt.Println("Banana price:", hashMap["banana"])
	fmt.Println("Orange price:", hashMap["orange"])

	hashMap["apple"] = 55
	fmt.Println("Apple price after updating:", hashMap["apple"])

	delete(hashMap, "banana")
	fmt.Println("Banana price after deleting:", hashMap["banana"])

}
