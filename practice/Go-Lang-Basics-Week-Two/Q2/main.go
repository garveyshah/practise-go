// Write a Go program that generates the first N terms of
// the Fibonacci sequence and stores them in a slice.
// Then, print the slice.
package main

import "fmt"

func main() {
	num := 1
	fmt.Println(Fibonacci(num))
}

func Fibonacci(num int) []int {
	ints := []int{}
	if num <= 1 {
		ints = append(ints, 1)
		return ints
	}
	ints = []int{1}
	a, b := 0, 1
	if num != 1 {
		for i := 1; i < num; i++ {
			c := a + b
			a = b
			b = c
			ints = append(ints, c)
		}
	}
	return ints
}
