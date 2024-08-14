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
	if num < 0 {
		return []int{}
		}
	
	result := []int{1}
	a, b := 0,1
	
	for i := 1; i <= num; i++ {
			c := a + b
			a = b
			b = c
		result = append(result, c)
	}
	return result
}
