package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	str := os.Args[1]
	a := os.Args[2]
	b := os.Args[3]

	result := SearchReplace(str, a, b)

	for _, char := range result {
		z01.PrintRune(char)
	}
	z01.PrintRune(10)
}

func SearchReplace(str string, a string, b string) string {
	newStr := ""
	for _, char := range str {
		if char == rune(a[0]) {
			char = rune(b[0])
		}
		newStr += string(char)
	}
	return newStr
}
