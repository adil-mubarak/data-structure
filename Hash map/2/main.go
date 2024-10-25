package main

import "fmt"

func main() {
	words := []string{"apple", "banana", "apple", "orange", "apple", "banana", "apple"}
	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}

	for word, count := range wordCount {
		fmt.Printf("%s: %d\n",word,count)
	}
}