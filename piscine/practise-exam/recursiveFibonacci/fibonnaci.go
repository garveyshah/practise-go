package main

import "fmt"

func Fibonacci(index int) int {
	if index == 1 || index == 2 {
		return 1
	} else {
		return Fibonacci(index-1) + Fibonacci(index-2)
	}
}
func main() {
	arg := 4
	fmt.Println(Fibonacci(arg))
}
