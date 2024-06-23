/*Day 2: Write a program to find the maximum and minimum in an array.*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Input the array of numbers separated by spaces:")

	reader := bufio.NewReader(os.Stdin)
	num, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(fmt.Errorf("error!! %w", err))
		return
	}

	num = strings.TrimSpace(num)
	nums := strings.Split(num, " ")

	numSlice := []int{}

	for _, char := range nums {
		num, err := CustomAtoi(char)
		if err != nil {
			fmt.Println(fmt.Errorf("error converting %s to integer: %w", char, err))
			return
		}
		numSlice = append(numSlice, num)
	}

	Max, Min := MaxAndMin(numSlice)
	fmt.Printf("maximum number is %d.\nMinimum number is %d.\n", Max, Min)
}

func MaxAndMin(s []int) (int, int) {
	BubbleSort(s)
	Max := s[len(s)-1]
	Min := s[0]
	return Max, Min
}

func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// Swap arr[j] andd arr[j+i]
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func CustomAtoi(s string) (int, error) {
	var result int
	for _, num := range s {
		if num < '0' || num > '9' {
			return 0, fmt.Errorf("invalid character: %c", num)
		}
		result = result*10 + int(num-'0')
	}
	return result, nil
}
