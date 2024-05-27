package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	for _, char := range os.Args[1] {
		z01.PrintRune(char)
	}

	z01.PrintRune('\n')
}
