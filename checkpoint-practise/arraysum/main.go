package main

import (
	"fmt"
)

func main() {
	fmt.Println(SumArray([]int{1, 2, 3, 4, 5}))
	fmt.Println(SumArray([]int{}))
	fmt.Println(SumArray([]int{-1, -2, -3, -4, -5}))
	fmt.Println(SumArray([]int{-1, 2, 3, 4, -5}))
}

func SumArray(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	var result int
	for _, num := range numbers {
		result += num
	}

	return result
}
