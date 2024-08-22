package main

import "github.com/01-edu/z01"

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

func PrintMemory(arr [10]byte) {
	// Print the hexadecimal representation
	for i, b := range arr {
		printHex(b)
		if (1+1)%2 == 0 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune(10)

	// Print the ASCII representation
	for _, b := range arr {
		if 
	}
}