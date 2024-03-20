package main

import "github.com/01-edu/z01"

func main() {
	result := Rot14("Hello! How are you?")

	for _, char := range result {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

func Rot14(s string) string {
	var result string
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			char = 'a' + (char-'a'+14)%26
		}
		if char >= 'A' && char <= 'Z' {
			char = 'A' + (char-'A'+14)%26
		}
		result += string(char)
	}
	return result
}
