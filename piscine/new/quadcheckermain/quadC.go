package main

import "github.com/01-edu/z01"

func QuadC(x, y int) {
	if x > 0 && y > 0 {
		for m := 1; m <= y; m++ {
			for b := 1; b <= x; b++ {
				if m == 1 && b == 1 {
					z01.PrintRune('A')
				} else if m == 1 && b > 1 && b < x {
					z01.PrintRune('B')
				} else if m == 1 && b == x {
					z01.PrintRune('A')
				} else if b == 1 && m > 1 && m < y {
					z01.PrintRune('B')
				} else if m > 1 && m < y && b > 1 && b < x {
					z01.PrintRune(' ')
				} else if m > 1 && m < y && b == x {
					z01.PrintRune('B')
				} else if m == y && b == 1 {
					z01.PrintRune('C')
				} else if m == y && b > 1 && b < x {
					z01.PrintRune('B')
				} else {
					z01.PrintRune('C')
				}
			}
			z01.PrintRune('\n')
		}
	}
}
