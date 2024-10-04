package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage : go run . <number>\nExample : go run . 10")
		return
	}

	num, err := conv.Atoi()
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	fib := compute.Fibonnaci(num)
	fmt.Printf("The %th Fibonnaci number is : %d", num, fib)
}
