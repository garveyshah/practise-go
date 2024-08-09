package main

import (
	"fmt"

	"project01/checkpoint-practise/itoa/itoa"
)

func main() {
	fmt.Println(itoa.Itoa(12345))
	fmt.Println(itoa.Itoa(0))
	fmt.Println(itoa.Itoa(-1234))
	fmt.Println(itoa.Itoa(987654321))
}
