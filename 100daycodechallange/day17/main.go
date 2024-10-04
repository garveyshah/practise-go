package main

import (
	"fmt"
	"os"

	"project01/100daycodechallange/day17/compute"
	"project01/100daycodechallange/day17/conv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage : go run . <number>\nExample : go run . 10")
		return
	}

	num, err := conv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	fib := compute.Fibonnaci(num)
	fmt.Printf("The %th Fibonnaci number is : %d\n", num, fib)
}
