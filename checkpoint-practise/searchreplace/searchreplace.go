package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	Print(SearchReplace(os.Args[1], os.Args[2], os.Args[3]))
}

func Print(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
	z01.PrintRune(10)
}

func SearchReplace(s, a, b string) string {
	var str string

	for _, char := range s {
		if string(char) == a {
			char = rune(b[0])
		}
		str += string(char)
	}
	return str
}
