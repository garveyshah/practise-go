/*
### 4. Generic Types
Modify the program to handle an array of strings instead of integers.
For example, given the input `["apple", "banana", "apple", "orange", "banana"]`, the program should output `["apple", "banana", "orange"]`.
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <slice of strings>")
		return
	}

	slice := os.Args[1:]
	fmt.Println(RemoveDuplicates(slice))
}

func RemoveDuplicates(s []string) []string {
	WordMap := make(map[string]bool)
	var temp []string

	for _, str := range s {
		WordMap[str] = true
	}

	for _, seen := range s {
		if _, ok := WordMap[seen]; ok {
			temp = append(temp, seen)
		}
		delete(WordMap, seen)
	}
	return temp
}
