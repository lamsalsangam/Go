package main

import "fmt"

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	var num1, num2 int
	fmt.Println("Enter the first number: ")
	fmt.Scan(&num1)
	fmt.Println("Enter the second number: ")
	fmt.Scan(&num2)

	result := gcd(num1, num2)

	fmt.Printf("The GCD of %d and %d is: %d", num1, num2, result)
}
