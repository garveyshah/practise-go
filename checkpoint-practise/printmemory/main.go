package main

import "github.com/01-edu/z01"

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

func PrintMemory(arr [10]byte) {
	// Print the hexadecimal representation
	for i, b := range arr {
		printHex(b)
		if (i+1)%2 == 0 {
			z01.PrintRune(' ')
		} 
	}
	z01.PrintRune(10)

	// Print the ASCII representation
	for _, b := range arr {
		if b >= 32 && b <= 126 {
			z01.PrintRune(rune(b))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune(10)
}

func printHex(b byte) {
	hexDigits := "0123456789abcdef"
	z01.PrintRune(rune(hexDigits[b/16]))
	z01.PrintRune(rune(hexDigits[b%16]))
	z01.PrintRune(' ')
	

}