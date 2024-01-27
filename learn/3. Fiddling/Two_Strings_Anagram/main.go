package main

import (
	"fmt"
	"slices"
	"strings"
)

func areAnagrams(str1, str2 string) bool {
	// Removing spaces and converting to lowercase for case-insensitive comparison
	str1 = strings.ReplaceAll(str1, " ", "")
	str2 = strings.ReplaceAll(str2, " ", "")
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	// Convert strings to slices of characters
	chars1 := strings.Split(str1, "")
	chars2 := strings.Split(str2, "")

	// Sort the slices
	slices.Sort(chars1)
	slices.Sort(chars2)

	// Check if the sorted slices are the same
	return strings.Join(chars1, "") == strings.Join(chars2, "")
}

func main() {
	string1 := "Listen"
	string2 := "Silent"

	result := areAnagrams(string1, string2)

	if result {
		fmt.Printf("%s and %s are anagrams.\n", string1, string2)
	} else {
		fmt.Printf("%s and %s are not anagrams.\n", string1, string2)
	}
}
