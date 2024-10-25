package main

import (
	"fmt"
	"unicode"
)

// how many letters

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

// reverse

func reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

//palindrome

func palindrome(pali string) bool {
	runes := []rune(pali)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true

}

// numbers of vowels

func vowels(vol string) int {
	count := 0
	for _, char := range vol {
		switch char {
		case 'a', 'i', 'o', 'e', 'u', 'A', 'E', 'I', 'O', 'U':
			count++
		}
	}
	return count
}

// letter shifting
func shift(shif string, n int) string {
	runes := []rune(shif)
	for i, char := range runes {
		if unicode.IsLetter(char) {
			if unicode.IsLower(char) {
				runes[i] = 'a' + (char - 'a' + rune(n)%26)
			} else {
				if unicode.IsUpper(char) {
					runes[i] = 'A' + (char - 'A' + rune(n)%26)
				}
			}
		}
	}
	return string(runes)
}

func backward(back string, n int) string {
	runes := []rune(back)
	for i, char := range runes {
		if unicode.IsLetter(char) {
			if unicode.IsLower(char) {
				runes[i] = 'a' + (char - 'a' - rune(n)%26)
			} else if unicode.IsUpper(char) {
				runes[i] = 'A' + (char - 'A' - rune(n)%26)
			}
		}
	}
	return string(runes)
}

func main() {
	str := "aaabbcdd"
	tub := look(str)
	fmt.Println(tub)

	rev := "anas"
	fmt.Println(reverse(rev))

	pali := "malayalam"
	fmt.Println(palindrome(pali))

	vol := "hello"
	fmt.Println(vowels(vol))

	shif := "adil"
	fmt.Println(shift(shif, 3))

	back := "dglo"
	fmt.Println(backward(back, 3))
}
