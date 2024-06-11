package main

import (
	"github.com/01-edu/z01"
)

func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}
	as := []int{500, 2}
	ReduceInt(as, mul)
	ReduceInt(as, sum)
	ReduceInt(as, div)
}

func ReduceInt(a []int, f func(int, int) int) {
	result := a[0]

	for i := 1; i < len(a); i++ {
		result = f(result, a[i])
	}

	if result < 0 {
		result = -result
	}
	var nums string
	if result == 0 {
		nums = "0"
	} else {
		for result > 0 {
			r := rune((result % 10) + '0')
			nums = string(r) + nums
			result /= 10
		}
	}

	for _, num := range nums {
		z01.PrintRune(num)
	}
	z01.PrintRune(10)
}
