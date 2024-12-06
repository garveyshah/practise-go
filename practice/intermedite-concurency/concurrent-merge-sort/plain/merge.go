package main

import (
	"fmt"
)

// Merge function that merges two sorted slices
func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0

	// Merge the two sorted slices
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			fmt.Println("left - ", result)
			i++
		} else {
			result = append(result, right[j])
			fmt.Println("right - ", result)

			j++
		}
	}

	// Append remaining elements from left or right slice
	result = append(result, left[i:]...)
	fmt.Println("left1 - ", result)

	result = append(result, right[j:]...)
	fmt.Println("right1 - ", result)

	return result
}

// MergeSort function to recursively split and sort the array
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("Unsorted array", arr)
	sortslice := mergeSort(arr)
	fmt.Println("Sorted array", sortslice)
}
