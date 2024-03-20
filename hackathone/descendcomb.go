package main

import "github.com/01-edu/z01"

func main() {
	DescendComb()
}

func DescendComb() {
	for i := 9; i >= 0; i-- {
		for j := i; j >= 0; j-- {
			if i != j {
				z01.PrintRune(rune(i) + '0')
				z01.PrintRune(rune(i) + '0')
				z01.PrintRune(',')
				z01.PrintRune(' ')
				if !(i == 0 && j == 0) {
					z01.PrintRune(rune(j) + '0')
					z01.PrintRune(rune(j) + '0')
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
