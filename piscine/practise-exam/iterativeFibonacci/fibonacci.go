package main

import "fmt"

func Fibonacci(index int) int {
	if index <= 1 {
		return 1
	}
	a, b := 0, 1
	if index != 1 {
		for i := 0; i <= index; i++ {
			c := a + b
			a = b
			b = c
		}
	}
	return index

}

func main() {
	arg := 4
	fmt.Println(Fibonacci(arg))
}
