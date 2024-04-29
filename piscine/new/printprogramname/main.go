package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	programName := args[0]
	programNameChars := []rune(programName)
	var programNameOnly []rune

	for i := len(programNameChars) - 1; i >= 0; i-- {
		if programNameChars[i] == '/' {
			programNameOnly = programNameChars[i+1:]
			break
		}
	}

	for _, char := range programNameOnly {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
