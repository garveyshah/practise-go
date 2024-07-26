package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 100

	fmt.Println(fizzBuzz(num))
}

func fizzBuzz(num int) []string {
	result := []string{}

	for i := 1; i <= num; i++ {
		switch {
		case i%3 == 0 && num%5 == 0:
			result = append(result, "FizzBuzz")
		case i%5 == 0:
			result = append(result, "Buzz")
		case i%3 == 0:
			result = append(result, "Fuzz")
		default:
			str := strconv.Itoa(i)
			result = append(result, str)
		}
	}
	return result
}
