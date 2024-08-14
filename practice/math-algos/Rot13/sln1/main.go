package main

import (
		"os"
		"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	args := os.Args[1]

	//result := []byte{}
	result := Rot13(args)

	// for i := 0; i < len(args); i++ {
	// 	char := args[i]
	// 	switch {
	// 	case char >= 'A' && char <= 'Z':
	// 		result = append(result, (char -'A' + 13)%26 + 'A')
	// 	case char >= 'a' && char <= 'z':
	// 		result = append(result, (char-'a'+13)%26 + 'a')
	// 	default:
	// 		result = append(result, char)
	// 	}
	// }

	for _, char := range string(result) {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

func Rot13(input string) string {
	var result []byte

	for i := 0; i < len(input); i++ {
		c := input[i]
		switch {
		case c >= 'A' && c <= 'Z':
			result = append(result, (c-'A'+13)%26+'A')
		case c >= 'a' && c <= 'z':
			result = append(result, (c-'a'+13)%26+'a')
		default:
			result = append(result, c)
		}
	}
	return string(result)
}
