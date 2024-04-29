package main

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}

	digit := n % 10
	z01.PrintRune(rune(digit) + '0')
	
	if n/10 != 0 {
		PrintNbrInOrder(n / 10)
	}
}

func main() {
	PrintNbrInOrder(321)
	PrintNbrInOrder(0)
	PrintNbrInOrder(321)
}
