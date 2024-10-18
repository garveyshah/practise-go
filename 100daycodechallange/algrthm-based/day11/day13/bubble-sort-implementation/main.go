package main

import (
	"fmt"
	"os"

	"project01/100daycodechallange/day13"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <arry to sorted>")
		return
	}
	var arr []int
	for _, num := range os.Args[1:] {
		num1, err := day13.CustomAtoi(num)
		if err != nil {
			fmt.Println("Error ", err)
		}
		arr = append(arr, num1)
	}

	sortedArr := BubbleSort(arr, func(a, b int) bool {
		return a > b
	})

	fmt.Println("Sorted Array: ", sortedArr)
}

func BubbleSort[T comparable](slice []T, less func(T, T) bool) []T {
	n := len(slice)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if less(slice[j], slice[j+1]) {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	return slice
}
