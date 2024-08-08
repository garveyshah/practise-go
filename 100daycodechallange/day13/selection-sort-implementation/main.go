/*
Problem 1: Selection Sort Implementation

- Given an array of integers, implement the selection sort algorithm to sort the array in ascending order.
*/
package main

import (
	"fmt"
	"os"

	"project01/100daycodechallange/day13"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run .  <array of intergers>")
		return
	}

	var slice []int

	for _, char := range os.Args[1:] {
		num, err := day13.CustomAtoi(char)
		if err != nil {
			fmt.Println("Error ", err)
		}
		slice = append(slice, num)
	}

	fmt.Println(SelectionSort(slice))
}

func SelectionSort(arr []int) []int {
	for i := 0; i <= len(arr)-1; i++ {
		for j := len(arr) - 1; j >= 0; j-- {
			if arr[i] > arr[j] {
				arr[j] = arr[i]
			} else  {
				arr[i] = arr[j]
			}
		}
	}
	return arr
}
