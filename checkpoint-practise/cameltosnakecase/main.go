package main

import (
	"github.com/01-edu/z01"
)

func CamelToSnakeCase(s string) string {
	// if s == "" {
	// 	return s
	// }

	// if !ValidCase(s) {
	// 	return s
	// }

	res := ""
	for i := 0; i < len(s); i++ {
		// If there if a Capital letter followed by a lowercase add an underscore before it and the letter not the first letter in the word
		if i != 0 && Up(rune(s[i])) && i+1 < len(s) && Low(rune(s[i+1])) {
			res += "_" + string(s[i])
			// res += string(s[i])
		} else if Low(rune(s[i])) || i == 0 && Up(rune(s[i])) {
			res += string(s[i])
		} else {
			return s
		}
	}
	return res
}

func ValidCase(str string) bool {
	for _, char := range str {
		if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') {
			return false
		}
	}
	return true
}

func Up(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func Low(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func Print(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
	z01.PrintRune(10)
}

func main() {
	Print(CamelToSnakeCase("HelloWorld"))
	Print(CamelToSnakeCase("helloWorld"))
	Print(CamelToSnakeCase("camelCase"))
	Print(CamelToSnakeCase("CAMELtoSnackCASE"))
	Print(CamelToSnakeCase("camelToSnakeCase"))
	Print(CamelToSnakeCase("hey2"))
}
