package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	var i int
	fmt.Println("Type the number you want to find the factorial of:")
	fmt.Scanln(&i)
	fmt.Println(fact(i))
}
