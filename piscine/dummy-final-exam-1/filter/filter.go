package main

import "fmt"

type Predicate func(int) bool

func Filtered(slice []int, fn Predicate) []int {
	filteredSlice := make([]int, 0)

	for _, num := range slice {
		if fn(num) {
			filteredSlice = append(filteredSlice, num)
		}
	}
	return filteredSlice
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	isEven := func(num int) bool {
		return num%2 == 0
	}

	evenNumbers := Filtered(numbers, isEven)

	fmt.Println("Even numbers:", evenNumbers)
}
