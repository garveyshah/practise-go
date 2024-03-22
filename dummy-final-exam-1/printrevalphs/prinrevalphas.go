package main

import "github.com/01-edu/z01"

func main() {
	for i := 'Z'; i >= 'A'; i-- {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')

	for i := 'z'; i >= 'a'; i-- {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
