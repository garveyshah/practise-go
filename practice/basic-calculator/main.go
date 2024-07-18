/*
	Write a program to perform basic arithmatic operations

(addition, subtration, multiplication, division) on two numbers.
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Enter an Arithmetic function.\n  ↓\n2 * 2\n\nEnsure you enclose the multiplication (*) operator in double quotation marks \"*\".")
		return
	}

	a, err := CustomAtoi(os.Args[1])
	if err != nil {
		fmt.Println("Err:", err)
		return
	}
	sign := os.Args[2]
	b, err := CustomAtoi(os.Args[3])
	if err != nil {
		fmt.Println("Error :-", err)
		return
	}

	switch {
	case sign == "/":
		if sign == "/" && ( b == 0) {
			fmt.Println("Make sure the divisor in the function is not 0")
		} else {
			fmt.Println(Div(a, b))
		}
	case sign == "*":
		fmt.Println(Mult(a, b))
	case sign == "+":
		fmt.Println(Add(a, b))
	case sign == "-":
		fmt.Println(Sub(a, b))
	default:
		fmt.Printf("Invalid Arithmetic Operator,\"%s\".\n\n Use: (\"*\") or (/) or (+) or (-).\n", sign)
		return
	}
}

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Mult(a, b int) int {
	return a * b
}

func Div(a, b int) int {
	return a / b
}

func CustomAtoi(s string) (int, error) {
	var result int
	sign := false

	if s[0] == '-' {
		sign = true
		s = s[1:]
	}

	for _, char := range s {
		if char < '0' || char > '9' {
			return 0, fmt.Errorf("error invalid character \"%v\"", string(char))
		}
		result = result*10 + int(char-'0')
	}
	if sign {
		result *= -1
	}
	return result, nil
}
