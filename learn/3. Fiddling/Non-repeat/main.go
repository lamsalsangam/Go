package main

import (
	"fmt"
)

func firstNonRepeatingChar(str string) rune {
	charCount := make(map[rune]int)

	// Count character occurrences
	for _, char := range str {
		charCount[char]++
	}

	// Find first character with a count of 1
	for _, char := range str {
		if charCount[char] == 1 {
			return char
		}
	}

	return 0 // No non-repeating characters found
}

func main() {
	str := "Avada ka dabra"
	firstChar := firstNonRepeatingChar(str)

	if firstChar != 0 {
		fmt.Printf("The first non-repeating character in '%s' is '%c'.\n", str, firstChar)
	} else {
		fmt.Println("All characters in the string repeat.")
	}
}
