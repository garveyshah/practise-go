package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	str1 := os.Args[1]
	str2 := os.Args[2]
	seen := make(map[rune]bool)
	var result string

	for _, runes := range str2 {
		seen[runes] = true
	}

	for _, val := range str1 {
		if seen[val] {
			result += string(val)
			delete(seen, val)
		}
	}
	for _, char := range result {
		z01.PrintRune(char)
	}
	z01.PrintRune(10)
}
