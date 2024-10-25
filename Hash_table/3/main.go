package main

import "fmt"

func main() {

	//sample text
	words := []string{"apple", "banana", "apple", "orange", "banana", "apple"}

	//hash map to store values
	wordCount := make(map[string]int)

	//count freaquncy of the number
	for _, word := range words {
		wordCount[word]++
	}

	//Output the freaquency count
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}
