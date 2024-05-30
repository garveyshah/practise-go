package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var str string
	if len(os.Args) != 3 {
		z01.PrintRune('\n')
		return
	}
	firstStr := os.Args[1]
	secondStr := os.Args[2]

	for _, letter := range firstStr {
		for _, char := range secondStr {
			if letter == char {
				str += string(letter)
				break
			} // else {
			// 	str = "\n"
			// }
		}
	}
	for _, char := range str {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
