package main

import (
	"fmt"
	"os"
	"project01/100daycodechallange/day14/funcs"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <NUMBER>\nExample: go run . 3222")
		return
	}

	sum, err := funcs.AddDigits(os.Args[1])
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Printf("The sum of the digits of %q is \"%d\"\n", os.Args[1], sum)
}
