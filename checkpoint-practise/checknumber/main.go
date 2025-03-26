package main

import "fmt"

func main() {
	fmt.Println(CheckNumber("Hello"))
	fmt.Println(CheckNumber("Hello1"))
	fmt.Println(CheckNumber(""))
}

func CheckNumber(input string) bool {
	if len(input) < 1 {
		return false
	}

	for _, char := range input {
		if char >= '0' && char <= '9' {
			return true
		}
	}
	return false
}
