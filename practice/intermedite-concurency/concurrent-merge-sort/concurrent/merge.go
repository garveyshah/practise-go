package main

import (
	"fmt"
	"sync"
)

// Merge function that merges two sorted slices
func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0

	// Merge the two sorted slices
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

// MergeSortConcurrent function to sort the array concurrently
func MergeSortConcurrent(arr []int, wg *sync.WaitGroup, ch chan []int) {
	defer wg.Done() // Decrease the counter when the function exits

	// Base case: a slice of lenth 1 or less is already sorted
	if len(arr) <= 1 {
		ch <- arr
		return
	}

	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]

	// Create channels for left and right halves
	leftCh := make(chan []int)
	rightChan := make(chan []int)

	var wg2 sync.WaitGroup
	wg2.Add(2)

	// Sort the left and right halves concurrently
	go func() {
		defer wg2.Done()
		MergeSortConcurrent(left, &wg2, leftCh)
	}()
	go func() {
		defer wg2.Done()
		MergeSortConcurrent(right, &wg2, rightChan)
	}()

	// Wait for the left and right goroutines to complete
	go func() {
		wg2.Wait()
		// Merge the sorted left and right halves
		merged := merge(<-leftCh, <-rightChan)
		ch <- merged
	}()
}

// Helper function to start concurrent merge sort
func concurrentMergeSort(arr []int) []int {
	ch := make(chan []int)
	var wg sync.WaitGroup
	wg.Add(1) // Wait for final mergeSortConcurrent function to finish

	go MergeSortConcurrent(arr, &wg, ch)

	wg.Wait() // wait for the mergeSortConcurrent goroutine to finish
	return <-ch
}

func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("Unsorted array:", arr)
	sortedArr := concurrentMergeSort(arr)
	fmt.Println("Sorted Array: ", sortedArr)
}
