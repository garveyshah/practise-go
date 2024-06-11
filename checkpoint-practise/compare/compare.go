package main

import (
	"github.com/01-edu/z01"
)

func main() {
	(Compare("Hello!", "Hello!"))
	(Compare("Salut!", "lut!"))
	(Compare("Ola!", "Ol"))
}

func Compare(a string, b string) {
	var result string
	if a == b {
		result = "0"
	}
	if a < b {
		result = "-1"
	}
	if a > b {
		result = "+1"
	}

	for _, num := range result {
		z01.PrintRune(num)
	}
	z01.PrintRune(10)
}
