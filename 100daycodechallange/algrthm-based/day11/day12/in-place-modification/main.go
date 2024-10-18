/*
### 6. In-place Modification
Write a Go program that removes duplicates from an array in-place, i.e., without using additional arrays or data structures. The program should modify the input array and return the new length of the array after duplicates are removed.
For example, given the input `[3, 3, 2, 1, 4, 2, 4]`, the program should modify the array to `[3, 2, 1, 4]` and return `4`.
*/

package main

import (
	"fmt"
	"os"

	handlingedgecases "project01/100daycodechallange/day12/handling-edge-cases"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage : go run . <array of ints>")
		return
	}

	var slice []int
	for _, char := range os.Args[1:] {
		num, err := handlingedgecases.CustomAtoi(char)
		if err != nil {
			fmt.Println("Error ", err)
			return
		}
		slice = append(slice, num)
	}

	fmt.Println(InPlaceMod(slice))
}

func InPlaceMod(s []int) int {
	var count int

	for i:= 0; i <= len(s)-1; i++ {
		count ++
		if count > 1 {
			s = RemoveIndex(s, s[i])
			
		} 
	}
	return len(s)
}

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}