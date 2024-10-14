package main

import "fmt"

func look(str string) string {
	runes := []rune(str)

	var result string

	for i := 0; i < len(runes); i++ {
		count := 1
		for i+1 < len(runes) && runes[i] == runes[i+1] {
			count++
			i++
		}
		result += fmt.Sprintf("%d%c", count, runes[i])
	}
	return result
}
func main() {
	str := "aaabbcdd"
	tub := look(str)
	fmt.Println(tub)
}
