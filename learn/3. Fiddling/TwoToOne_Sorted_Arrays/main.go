package main

import (
	"fmt"
	"slices"
)

func mergedSortedArrays(arr1 []int, arr2 []int) []int {
	// merged := make([]int, 0, len(arr1)+len(arr2))
	merged := append([]int{}, arr1...)
	merged = append(merged, arr2...)
	// merged := append(append([]int{}, arr1...), arr2...)

	// merged := make([]int, len(arr1)+len(arr2))
	// i := 0
	// for _, val := range arr1 {
	// 	merged[i] = val
	// 	i++
	// }
	// for _, val := range arr2 {
	// 	merged[i] = val
	// 	i++
	// }

	slices.Sort(merged)

	return merged
}

func main() {
	arr1 := []int{1, 3, 5}
	arr2 := []int{8, 4, 6, 2}

	mergedArr := mergedSortedArrays(arr1, arr2)
	fmt.Println(mergedArr)
}
