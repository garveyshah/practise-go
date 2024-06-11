package main

import (
	//"os"

	"github.com/01-edu/z01"
)

// func main() {
// 	if len(os.Args) != 2 {
// 		return
// 	}
// 	alpha := os.Args[1]

// 	Mirror := AlpaMirror(alpha)

// 	for _, char := range Mirror {
// 		z01.PrintRune(char)
// 	}
// 	z01.PrintRune(10)
// }

// // Method 2 : Mapping
// func AlpaMirror(s string) string {
// 	alpha := map[rune]rune{'a': 'z', 'b': 'y', 'c': 'x', 'd': 'w', 'e': 'v', 'f': 'u', 'g': 't', 'h': 's', 'i': 'r', 'j': 'q', 'k': 'p', 'l': 'o', 'm': 'n', 'n': 'm', 'o': 'l', 'p': 'k', 'q': 'j', 'r': 'i', 's': 'h', 't': 'g', 'u': 'f', 'v': 'e', 'w': 'd', 'x': 'c', 'y': 'b', 'z': 'a', 'A': 'Z', 'B': 'Y', 'C': 'X', 'D': 'W', 'E': 'V', 'F': 'U', 'G': 'T', 'H': 'S', 'I': 'R', 'J': 'Q', 'K': 'P', 'L': 'O', 'M': 'N', 'N': 'M', 'O': 'L', 'P': 'K', 'Q': 'J', 'R': 'I', 'S': 'H', 'T': 'G', 'U': 'F', 'V': 'E', 'W': 'D', 'X': 'C', 'Y': 'B', 'Z': 'A'}
// 	Mirror := ""
// 	for _, char := range s {
// 		value, ok := alpha[char]
// 		if ok {
// 			Mirror += string(value)
// 		} else {
// 			Mirror += string(char)
// 		}
// 	}
// 	return Mirror
// }

func main() {
	for j := 'A'; j <= 'Z'; j++ {
		paired := 'Z' - (j - 'A')
		z01.PrintRune('\'')
		z01.PrintRune(j)
		z01.PrintRune('\'')
		z01.PrintRune(':')
		if paired != j {
			z01.PrintRune('\'')
			z01.PrintRune(paired)
			z01.PrintRune('\'')
			z01.PrintRune(',')
		}

	}
	z01.PrintRune(10)
}
