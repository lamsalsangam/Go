// package main

// import (
// 	"fmt"
// 	"slices"
// )

// func min(slice []int) int {
// 	if len(slice) == 0 {
// 		return 0 // or any default value, depending on your requirements
// 	}

// 	minimum := slice[0]
// 	for _, value := range slice {
// 		if value < minimum {
// 			minimum = value
// 		}
// 	}
// 	return minimum
// }

// func main() {
// 	arr1 := []int{23, 26, 1, 5, 7, 17, 2, 5, 50}
// 	slices.Sort(arr1)
// 	minimum := min(arr1)
// 	fmt.Println("Minimum:", minimum)
// }

// package main

// import (
// 	"fmt"
// 	"slices"
// )

// func main() {
// 	arr1 := []int{23, 26, 1, 5, 7, 17, 2, 50}
// 	slices.Sort(arr1)
// 	fmt.Println(slices.Min(arr1))
// 	fmt.Println(slices.Max(arr1))
// }

// package main

// import (
// 	"fmt"
// 	"slices"
// )

// func min(slice []int) int {
// 	if len(slice) == 0 {
// 		return 0
// 	}

// 	minimum := slice[0]
// 	for _, v := range slice {
// 		if v < minimum {
// 			minimum = v
// 		}
// 	}
// 	return minimum
// }

// func max(slice []int) int {
// 	if len(slice) == 0 {
// 		return 0
// 	}
// 	maximum := slice[0]
// 	for _, v := range slice {
// 		if v > maximum {
// 			maximum = v
// 		}
// 	}
// 	return maximum
// }

// func main() {
// 	arr1 := []int{23, 26, 1, 5, 7, 17, 2, 50}
// 	slices.Sort(arr1)
// 	minimum := min(arr1)
// 	fmt.Println("Minimum:", minimum)
// 	maximum := max(arr1)
// 	fmt.Println("Maximum:", maximum)
// }

package main

import (
	"fmt"
)

func min(slice []int, ch chan int) {
	// Find the minimum in the slice and send it to the channel
	minimum := slice[0]
	for _, v := range slice {
		if v < minimum {
			minimum = v
		}
	}
	ch <- minimum
}

func max(slice []int, ch chan int) {
	// Find the maximum in the slice and send it to the channel
	maximum := slice[0]
	for _, v := range slice {
		if v > maximum {
			maximum = v
		}
	}
	ch <- maximum
}

func main() {
	arr1 := []int{23, 26, 1, 5, 7, 17, 2, 50}

	// Create channels to receive min and max values
	minCh := make(chan int)
	maxCh := make(chan int)

	// Launch goroutines to find min and max concurrently
	go min(arr1, minCh)
	go max(arr1, maxCh)

	// Receive and print minimum and maximum values
	minimum := <-minCh
	maximum := <-maxCh

	fmt.Println("Minimum:", minimum) 
	fmt.Println("Maximum:", maximum)
}
