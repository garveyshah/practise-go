package main

import (
	"fmt"
	"os"

	"project01/100daycodechallange/day15/funcs"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage : go run . <Year>\nExample : go run . 1990")
		return
	}

	year, err := funcs.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Err : ", err)
		return
	}

	if funcs.IsLeapYear(year) {
		fmt.Printf("\"%d\" is a leap year.", year)
	} else {
		fmt.Printf("\"%d\" is NOT a leap year", year)
	}
}
