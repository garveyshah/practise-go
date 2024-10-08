package main

import (
	"fmt"
	"os"
	"project01/100daycodechallange/day19/calc"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage : go run . <number>\nExample : go run . 90")
		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Err : ", err)
		return
	}

	exp, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Err : ", err)
		return
	}

	power := calc.Power(num, exp)

	fmt.Printf("The power of %d  raised to the number %d is %d", num, exp, power)
}
