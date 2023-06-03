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
func maxMessages(thresh float64) int {
	// ?
	totalCost := 0.0
	for i := 0; ; i++ {
		totalCost += 1.0 + (0.01 * float64(i))
		if totalCost > thresh {
			return i
		}
	}
}

// don't edit below this line

func test(thresh float64) {
	fmt.Printf("Threshold: %.2f\n", thresh)
	max := maxMessages(thresh)
	fmt.Printf("Maximum messages that can be sent: = %v\n", max)
	fmt.Println("===============================================================")
}

func main() {
	test(10.00)
	test(20.00)
	test(30.00)
	test(40.00)
	test(50.00)
}
