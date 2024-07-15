package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	num, _ := customAtoi(os.Args[1])

	result := IsPowerOf2(num)

	Printer(result)
}

func Printer(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
	z01.PrintRune(10)
}

func 

func IsPowerOf2(n int) string {
	if n > 0 && n&(n-1) == 0 {
		return "true"
	}
	return "false"
}

func customAtoi(s string) (int, error) {
	if s == "0" {
		return 0, nil
	}
	var err error
	sign := false
	result := 0
	if s[0] == '-' {
		sign = true
		s = string(s[1:])
	}

	for _, char := range s {
		if char < '0' || char > '9' {
			return 0, err
		}
		result = result*10 + int(char-'0')
	}
	if sign {
		result *= -1
	}
	return result, nil
}
