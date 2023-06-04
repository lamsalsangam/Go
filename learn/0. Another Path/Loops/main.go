package main

import (
	"fmt"
)

// for INITIAL; CONDITION; AFTER{
// 	// do something
// }
// for i := 0; i < 10; i++ {
// 	fmt.Println(i)
// }
// // Prints 0 through 9

// ===============================

// func bulkSend(numMessages int) float64 {
// 	// ?
// 	totalCost := 0.0
// 	for i := 0; i < numMessages; i++ {
// 		totalCost += 1.0 + (0.01 * float64(i))
// 	}
// 	return totalCost
// }

// // don't edit below this line

// func test(numMessages int) {
// 	fmt.Printf("Sending %v messages\n", numMessages)
// 	cost := bulkSend(numMessages)
// 	fmt.Printf("Bulk send complete! Cost = %.2f\n", cost)
// 	fmt.Println("===============================================================")
// }

// func main() {
// 	test(10)
// 	test(20)
// 	test(30)
// 	test(40)
// 	test(50)
// }

// ################################################################
// Omitting Conditions
// func maxMessages(thresh float64) int {
// 	// ?
// 	totalCost := 0.0
// 	for i := 0; ; i++ {
// 		totalCost += 1.0 + (0.01 * float64(i))
// 		if totalCost > thresh {
// 			return i
// 		}
// 	}
// }

// // don't edit below this line

// func test(thresh float64) {
// 	fmt.Printf("Threshold: %.2f\n", thresh)
// 	max := maxMessages(thresh)
// 	fmt.Printf("Maximum messages that can be sent: = %v\n", max)
// 	fmt.Println("===============================================================")
// }

// func main() {
// 	test(10.00)
// 	test(20.00)
// 	test(30.00)
// 	test(40.00)
// 	test(50.00)
// }

// ################################################################

//	for CONDITION {
//		// do some stuff while CONDITION is true
//	  }

// ===============================

// func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
// 	actualCostInPennies := 1.0
// 	maxMessagesToSend := 0
// 	for actualCostInPennies <= float64(maxCostInPennies) {
// 		maxMessagesToSend++
// 		actualCostInPennies *= costMultiplier
// 	}
// 	return maxMessagesToSend
// }

// // don't touch below this line

// func test(costMultiplier float64, maxCostInPennies int) {
// 	maxMessagesToSend := getMaxMessagesToSend(costMultiplier, maxCostInPennies)
// 	fmt.Printf("Multiplier: %v\n",
// 		costMultiplier,
// 	)
// 	fmt.Printf("Max cost: %v\n",
// 		maxCostInPennies,
// 	)
// 	fmt.Printf("Max messages you can send: %v\n",
// 		maxMessagesToSend,
// 	)
// 	fmt.Println("====================================")
// }

// func main() {
// 	test(1.1, 5)
// 	test(1.3, 10)
// 	test(1.35, 25)
// }

// ################################################################

func printPrimes(max int) {
	// ?
	for n := 2; n < max+1; n++ {
		if n == 2 {
			fmt.Println(n)
			continue
		}
		if n%2 == 0 {
			continue
		}
		isPrime := true
		for i := 3; i*i < n+1; i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}
		if !isPrime {
			continue
		}
		fmt.Println(n)
	}
}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}
