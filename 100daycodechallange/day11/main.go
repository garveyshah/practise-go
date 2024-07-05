/*
Day 11: Implement a function to perform binary search on a sorted array.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	binarysearch "project01/100daycodechallange/day11/binarySearch"
)

func main() {
	fmt.Println("enter the string of numbers separated by spaces")

	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error!!...reading input")
	}

	str = strings.TrimSpace(str)
	str1 := strings.Split(str, "")
	var Array []int

	for _, char := range str1 {
		num, err := binarysearch.CustomAtoi(char)
		if err != nil {
			fmt.Println("Error:..", err)
			return
		}
		Array = append(Array, num)
	}
}
