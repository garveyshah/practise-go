package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	Print(Mirror(os.Args[1]))
}

func Mirror(s string) string {
	var Mirror string
	for i := len(s) - 1; i >= 0; i-- {
		Mirror += string(s[i])
	}
	return Mirror
}

func Print(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
	z01.PrintRune(10)
}
