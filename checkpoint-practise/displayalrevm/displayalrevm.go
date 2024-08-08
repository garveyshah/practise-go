package main

import "github.com/01-edu/z01"

func main() {
	str := "abcdefghijklmnopqrstuvwxyz"
	for i := len(str) - 1; i >= 0; i-- {
		char := str[i]
		if i%2 == 0 {
			z01.PrintRune(rune(char - 32))
			continue
		}
		z01.PrintRune(rune(char))
	}
	z01.PrintRune('\n')
}
