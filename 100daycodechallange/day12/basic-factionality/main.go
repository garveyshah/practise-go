// Basic Functionality
// Write a Go program that takes an array of integers and returns a new array with all duplicate elements removed. For example, given the input `[1, 2, 2, 3, 4, 4, 5]`, the program should output `[1, 2, 3, 4, 5]`.
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <array of numbers>")
		return
	}
	slice := SplitStrings(os.Args[1])

	var arr []int
	for _, char := range slice {
		num, err := CustomAtoi(string(char))
		if err != nil {
			fmt.Println("Error!! ", err)
			return
		}
		arr = append(arr, num)
	}
	fmt.Println(arr)
}

// func RemoveDuplicates(arr []int) []int {
// }

func CustomAtoi(s string) (int, error) {
	var result int
	var isNeg bool

	if s[0] == '-' {
		isNeg = true
		s = s[1:]
	}

	for _, char := range s {
		if char < '0' || char > '9' {
			return 0, fmt.Errorf("invalid character \"%v\"", string(char))
		}
		result = result*10 + int(char-'0')
	}

	if isNeg {
		result *= -1
	}
	return result, nil
}

func SplitStrings(s string) []string {
	var result []string
	var cup string

	for i:=0; i <= len(s)-1; i {
		
		if char == ' ' || i == len(s) {
			result = append(result, cup)
			cup = ""
		}
		cup += string(char)
	}
	return result
}
