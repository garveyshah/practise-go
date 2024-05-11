package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Arguments less than 2")
		return
	}

	args := os.Args[1:]

	num, err := strconv.Atoi(args[0])
	if err != nil {
		log.Println("Error")
	}
	result := Fibbs(num)

	fmt.Println(result)
}

func Fibbs(num int) []int {
	result := []int{1}
	for a := 1; a <= num; a++ {
		for b := 1; b <= num; b++ {
			c := a + b
			b = c
			a = b
			if a + b == c {
			result = append(result, c)
			}
		}
	}
	return result
}
