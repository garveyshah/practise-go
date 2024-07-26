package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	Value1 := os.Args[1]
	operator := (os.Args[2])
	Value2 := os.Args[3]
	result := 0

	num1, err := StringToInt(Value1)
	if err != nil {
		return
	}

	num2, err := StringToInt(Value2)
	if err != nil {
		return
	}
	if (num1 >= 9223372036854775807) || (num2 >= 9223372036854775807) || (num1 <= int(-9223372036854775807)) || (num2 <= -9223372036854775807) {
		return
	}
	if operator == "-" {
		result = num1 - num2
	} else if operator == "+" {
		result = num1 + num2
	} else if operator == "/" {
		if num1 == 0 || num2 == 0 {
			fmt.Println("No division by 0")
			return
		} else {
			result = num1 / num2
		}
	} else if operator == "%" {
		if num1 == 0 || num2 == 0 {
			fmt.Println("No modulo by 0")
			return
		} else {
			result = num1 % num2
		}
	} else if operator == "*" {
		result = num1 * num2
		//fmt.Println(result)
	} else {
		return
	}
	fmt.Println(result)
}

func StringToInt(s string) (int, error) {
	var result int
	var err error
	sigh := 1
	if s[0] == '-' {
		s = s[1:]
		sigh = -1
	}
	for _, char := range s {
		if char < '0' || char > '9' {
			return 0, err
		}

		result = result*10 + int(char-'0')
	}

	result = sigh * result
	return result, err
}
