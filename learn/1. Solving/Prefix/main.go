package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0] // Initialize prefix as the first string in the array

	// Iterate through each string in the array
	for i := 1; i < len(strs); i++ {
		// Keep reducing the prefix until it becomes the common prefix
		for !strings.HasPrefix(strs[i], prefix) {
			prefix = prefix[:len(prefix)-1]

			// If prefix becomes empty, there is no common prefix
			if prefix == "" {
				return ""
			}
		}
	}

	return prefix
}

func main() {
	strs := []string{"flower", "flow", "flight"}

	result := longestCommonPrefix(strs)
	fmt.Printf("Longest common prefix: %s\n", result)
}
