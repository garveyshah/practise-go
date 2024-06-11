package main

import (
	"os"

	if len(os.Args) != 3 {
		return
	}

	Value1 := os.Args[1]
	operator := rune(os.Args[2])
	Value2 := os.Args[3]
	result := 0
	if operator == '-' {
		result = Value1 - Value2
	}
	if result == '+' {
		result = Value1 + Value
	}


)