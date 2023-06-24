package main

import (
	"fmt"
	"strings"
)

func generateCombinations(chars string) []string {
	var combinations []string
	length := len(chars)

	for i := 0; i < (1 << length); i++ {
		var combination strings.Builder

		for j := 0; j < length; j++ {
			if (i & (1 << j)) > 0 {
				combination.WriteByte(chars[j])
			}
		}

		combinations = append(combinations, combination.String())
	}

	return combinations
}

func main() {
	input := "abc"
	combinations := generateCombinations(input)

	fmt.Println("All possible combinations:")
	for _, combination := range combinations {
		fmt.Println(combination)
	}
}
