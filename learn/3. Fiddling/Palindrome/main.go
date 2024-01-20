package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	// Convert the string to lowercase and remove spaces
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))

	// Check if the string is a palindrome
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	input := "A man a plan a canal Panama"
	if isPalindrome(input) {
		fmt.Printf("%s is a palindrome.\n", input)
	} else {
		fmt.Printf("%s is not a palindrome.\n", input)
	}
}
