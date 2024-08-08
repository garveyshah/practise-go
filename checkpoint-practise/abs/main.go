package main

import (
	"fmt"
	"os"

	"project01/100daycodechallange/day13"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	num, _ := day13.CustomAtoi(os.Args[1])
	 
	fmt.Println(num)
	fmt.Println(IsPositive(num))
	fmt.Println(Abs(num))
}

func IsPositive(n int) bool {
	return n >= 0
}

func Abs(n int) int {
	if n < 0 {
		n *= -1
	}

	return n
}
