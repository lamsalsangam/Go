package main

import "fmt"

func isPerfectNumber(num int) bool {
	if num <= 0 {
		return false
	}

	divisorsSum := 0

	// Find the proper divisors and sum them
	for i := 1; i*i <= num; i++ {
		if num%i == 0 {
			divisorsSum += i

			// Add the corresponding divisor if it is not the number itself
			if i != num/i && i != 1 {
				divisorsSum += num / i
			}
		}
	}

	// Check if the sum of divisors is equal to the number
	return divisorsSum == num
}

func main() {
	num := 28

	if isPerfectNumber(num) {
		fmt.Printf("%d is a perfect number.\n", num)
	} else {
		fmt.Printf("%d is not a perfect number.\n", num)
	}
}
