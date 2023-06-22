package main

import (
	"fmt"
	"sort"
)

func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	var size, target int
	fmt.Println("Enter the size of the array:")
	fmt.Scan(&size)

	arr := make([]int, size)
	fmt.Println("Enter the values of the array:")
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}
	sort.Ints(arr)

	fmt.Println("Enter the target values:")
	fmt.Scanln(&target)

	index := binarySearch(arr, target)
	fmt.Println("Index of", target, "is", index)
}
