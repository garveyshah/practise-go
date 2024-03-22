package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var arry []int
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	for n > 0 {
		arry = append(arry, n%10)
		n = n / 10
	}
	for i := 0; i < len(arry); i++ {
		for j := i + 1; j < len(arry); j++ {
			if arry[i] > arry[j] {
				arry[i], arry[j] = arry[j], arry[i]
			}
		}
	}
	for _, digit := range arry {
		z01.PrintRune(rune(digit) + '0')
	}
}
