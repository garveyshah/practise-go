/*
## 2. Handling Edge Cases
Write a Go program to remove duplicates from an array that contains only one element, an empty array, and an array where all elements are the same. For example:
- Input: `[7]` → Output: `[7]`
- Input: `[]` → Output: `[]`
- Input: `[5, 5, 5, 5, 5]` → Output: `[5]`
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println([]int{})
	}

	slice := os.Args[1:]
	fmt.Println("input Arr - ", slice)

	numsArr := []int{}
	for _, char := range slice {
		num, err := CustomAtoi(char)
		if err != nil {
			fmt.Println("Error ", err)
			return
		}
		numsArr = append(numsArr, num)
	}
	fmt.Println("Cleaned arr", RemoveDuplicates(numsArr))
}

func RemoveDuplicates(arr []int) []int {
	numsM := make(map[int]bool)
	Clean := []int{}

	if len(arr) == 0 || len(arr) == 1 {
		return arr
	}

	for _, num := range arr {
		numsM[num] = true
	}

	for _, seen := range arr {
		if _, ok := numsM[seen]; ok {
			Clean = append(Clean, seen)
		}
		delete(numsM, seen)
	}
	return Clean
}

func CustomAtoi(s string) (int, error) {
	var result int
	var isNeg bool

	if s[0] == ' ' {
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
