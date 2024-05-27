package main

import "github.com/01-edu/z01"

func main() {
	str := "Hello World"

	for _, char := range str {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
