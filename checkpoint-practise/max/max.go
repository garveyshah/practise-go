package main

import "github.com/01-edu/z01"

func main() {
	arr := []int{23, 123, 1, 11, 55, 93}
	max := Max(arr)

	if max < 0 {
		max = -max
	}

	var str string

	if max == 0 {
		str = "0"
	} else {
		for max > 0 {
			r := (rune(max%10) + '0')
			str = string(r) + str
			max /= 10
		}
	}
	for _, char := range str {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

func Max(arr []int) int {
	num := 0

	for _, char := range arr {
		if char > num {
			num = char
		}
	}
	return num
}
