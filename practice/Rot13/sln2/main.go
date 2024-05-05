package main

import (
	"fmt"
	"github.com/01-edu/z01"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	args := os.Args[1]

	fmt.Println(args)
	result := []byte{}

	alphLib := map[byte]byte{'A': 'N', 'B': 'O', 'C': 'P', 'D': 'Q', 'E': 'R', 'F': 'S', 'G': 'T', 'H': 'U', 'I': 'V', 'J': 'W', 'K': 'X', 'L': 'Y', 'M': 'Z', 'N': 'A', 'O': 'B', 'P': 'C', 'Q': 'D', 'R': 'E', 'S': 'F', 'T': 'G', 'U': 'H', 'V': 'I', 'W': 'J', 'X': 'K', 'Y': 'L', 'Z': 'M', 'a': 'n', 'b': 'o', 'c': 'p', 'd': 'q', 'e': 'r', 'f': 's', 'g': 't', 'h': 'u', 'i': 'v', 'j': 'w', 'k': 'x', 'l': 'y', 'm': 'z', 'n': 'a', 'o': 'b', 'p': 'c', 'q': 'd', 'r': 'e', 's': 'f', 't': 'g', 'u': 'h', 'v': 'i', 'w': 'j', 'x': 'k', 'y': 'l', 'z': 'm'}

	for i := 0; i < len(args); i++ {
		char1 := args[i]
		if val, ok := alphLib[char1]; ok {
			result = append(result, val)
		} else {
			result = append(result, char1)
		}
	}
	Result2 := string(result)
	
	for _, char := range Result2 {
		z01.PrintRune(rune(char))
	}
	z01.PrintRune('\n')
}
