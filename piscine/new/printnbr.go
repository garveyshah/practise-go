package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	a := 1
	if n == 0 {
		z01.PrintRune('0')
	}
	if n < 0 {
		a = -1
		z01.PrintRune('-')
	}
	if n != 0 {
		b := (n / 10) * a
		if b != 0 {
			PrintNbr(b)
		}
		c := (n % 10 * a) + '0'
		runeC := rune(c)
		z01.PrintRune(runeC)
	}
}
