package main

import (
	"fmt"
)

func main() {
	fmt.Println(ToUpper("Hello! How are you?"))
}

func ToUpper(s string) string {
	var new string
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			char = char - 32
		}
		new += string(char)
	}
	return new
}
