package main

import "fmt"

// binarySearch performs binary search on a sorted array.
func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == target {
			return mid // target found
		} else if arr[mid] < target {
			low = mid + 1 // search in the right half
		} else {
			high = mid - 1 // search in the left half
		}
	}

	return -1 // target not found
}

func main() {
	// Example usage:
	sortedArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 7

	index := binarySearch(sortedArray, target)

	if index != -1 {
		fmt.Printf("Target %d found at index %d\n", target, index)
	} else {
		fmt.Printf("Target %d not found in the array\n", target)
	}
}
