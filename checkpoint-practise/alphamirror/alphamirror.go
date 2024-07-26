package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Method 1: Algorithm
func main() {
	if len(os.Args) != 2 {
		return
	}

	alpha := os.Args[1]
	Mirror := ""

	for _, char := range alpha {
		switch {
		case char >= 'A' && char <= 'Z':
			Mirror += string('Z' - (char - 'A'))
		case char >= 'a' && char <= 'z':
			Mirror += string('z' - (char - 'a'))
		default:
			Mirror += string(char)
		}
	}

	for _, char := range Mirror {
		z01.PrintRune(char)
	}
	z01.PrintRune(10)
}
