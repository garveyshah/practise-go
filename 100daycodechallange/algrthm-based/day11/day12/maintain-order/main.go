/*## 5. Maintaining Ordera Go program to remove duplicates from an array while preserving the order of the first occurrence of each element
Write .
For example, given the input `[4, 5, 6, 4, 5, 7]`, the program should output `[4, 5, 6, 7]`.
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Println("Usage: go run . <arry  of numbers>")
		return
	}

	slice := os.Args[1:]
	if len(slice) == 0 {
		fmt.Println([]int{})
		return
	}

	num := []int{}
	for _, char := range slice {
		char, err := CustomAtoi(char)
		if err != nil {
			fmt.Println("Error ", err)
		}
		num = append(num, char)
	}

	fmt.Println(RemoveDups(num))
}

func RemoveDups(arr []int) []int {
	NumsMap := make(map[int]bool)
	result := []int{}

	for _, char := range arr {
		NumsMap[char] = true
	}

	for _, char := range arr {
		if _, ok := NumsMap[char]; ok {
			result = append(result, char)
		}
		delete(NumsMap, char)
	}
	return result
}

func CustomAtoi(s string) (int, error) {
	result := 0
	isNeg := false

	if s[0] == ' ' {
		isNeg = true
		s = s[1:]
	}

	for _, char := range s {
		if char < '0' || char > '9' {
			return 0, fmt.Errorf("invalid char \"%v\"", string(char))
		}
		result = result*10 + int(char-'0')
	}
	if isNeg {
		result *= -1
	}
	return result, nil
}
