package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Occurrence struct {
	Counts map[rune]int
}

func CountCharacters(s string) Occurrence {
	counts := make(map[rune]int)

	// Iterate over each character in the string
	for _, char := range s {
		// Increment the count for the character
		counts[char]++
	}

	occurrence := Occurrence{
		Counts: counts,
	}

	return occurrence
}

func main() {
	// Prompt the user to enter a string
	fmt.Println("Enter a string:")

	// Read the entire line of input
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// Remove the newline characters from the input
	input = strings.TrimSpace(input)

	occurrence := CountCharacters(input)

	// Print the character counts
	for char, count := range occurrence.Counts {
		fmt.Printf("%c: %d\n", char, count)
	}
}
