package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please parse an argument")
		return
	}

	result := 0

	for _, char := range os.Args[1] {
		digit := int(char - '0')
		if digit < 0 || digit > 9 {
			fmt.Println("ERROR: cannot convert to roman digit.")
			return
		}

		result = result*10 + digit
	}

	if result == 0 || result == 4000 {
		fmt.Println("ERROR: cannot convert to roman digit.")
		return
	}

	RomanNums := map[int]string{
		1: "I", 4: "IV", 5: "V",
		9: "IX", 10: "X", 40: "XL",
		50: "L", 90: "XC", 100: "C",
		400: "CD", 500: "D", 900: "CM",
		1000: "M",
	}

	// Slice to maintain the order of the Roman numeral values
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	final := ""

	for _, value := range values {
		for result >= value {
			final += RomanNums[value]
			result -= value
		}
	}
	fmt.Printf("%s\n", final)
}
