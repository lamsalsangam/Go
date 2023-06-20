package main

import (
	"fmt"
)

func palindrome(value string) string {
	reverse := ""
	for _, v := range value {
		reverse = string(v) + reverse
	}
	return reverse
}

func main() {
	var val string
	fmt.Println("What do you want to find the palindrome of:")
	fmt.Scanln(&val)
	reverse := palindrome(val)
	if val != reverse {
		fmt.Println("It is not a palindrome")
		return
	}
	fmt.Println("It is palindrome")
}
