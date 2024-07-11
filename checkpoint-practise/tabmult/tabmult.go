package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune(10)
		return
	}
	num, _ := customAtoi(os.Args[1])
	product := 0

	for i := 1; i <= 9; i++ {
		product = i * num
		Printer(customItoa(i) + " x " + customItoa(num) + " = " + customItoa(product))
	}
}

func Printer(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
	z01.PrintRune(10)
}

func customAtoi(s string) (int, error) {
	var err error
	var result int
	for _, char := range s {
		if char < '0' || char > '9' {
			return 0, err
		}
		result = result*10 + int(char-'0')
	}
	return result, nil
}

func customItoa(num int) string {
	sign := 1
	if num == 0 {
		return "0"
	}
	if num < 0 {
		sign *= -1
		num *= -1
	}
	result := ""
	for num > 0 {
		r := num % 10
		result = string(r+('0')) + result
		num /= 10
	}
	if sign == -1 {
		result = "-" + result
	}
	return result
}
