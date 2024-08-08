package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	params := os.Args[1:]
	for i := len(params) - 1; i >= 0; i-- {
		Print(params[i])
	}
}

func Print(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
	z01.PrintRune(10)
}
