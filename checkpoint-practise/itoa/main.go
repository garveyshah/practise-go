package main

import (
	"fmt"
)

func main() {
	fmt.Println(Itoa(12345))
	fmt.Println(Itoa(0))
	fmt.Println(Itoa(-1234))
	fmt.Println(Itoa(987654321))
}

func Itoa(n int) string {
	var result string
	var IsNeg bool

	if n < 0 {
		IsNeg = true
		n *= -1
	}

	if n == 0 {
		return "0"
	}

	for n > 0 {

		r := n%10 +'0'
		result = string(r) + result
		n = n / 10
	}
	if IsNeg {
		result = "-" + result
	}
	return result
}
