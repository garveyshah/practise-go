package main

import "fmt"

func main() {
	fmt.Println(PopInt([]int{6, 7, 8, 9}))
}

func PopInt(ints []int) []int {
	if len(ints) == 0 {
		return []int{}
	}

	var new []int
	new = append(new, ints[:len(ints)-1]...)
	return new
}
