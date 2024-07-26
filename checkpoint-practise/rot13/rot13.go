package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	plainText := os.Args[1]
	cypherText := ""
	for _, char := range plainText {
		switch {
		case char >= 'A' && char <= 'Z':
			cypherText += string((char+13-'A')%26 + 'A')
		case char >= 'a' && char <= 'z':
			cypherText += string((char+13-'a')%26 + 'a')
		default:
			cypherText += string(char)
		}
	}
	for _, char := range cypherText {
		z01.PrintRune(char)
	}
	z01.PrintRune(10)
}
