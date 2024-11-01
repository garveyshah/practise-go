package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// var str string
	if len(os.Args) != 3 {
		z01.PrintRune('\n')
		return
	}
	firstStr := os.Args[1]
	secondStr := os.Args[2]

	fmt.Println(Wmatch(firstStr, secondStr))

	// for _, letter := range firstStr {
	// 	for _, char := range secondStr {
	// 		if letter == char {
	// 			str += string(letter)
	// 			break
	// 		} // else {
	// 		// 	str = "\n"
	// 		// }
	// 	}
	// }
	// for _, char := range str {
	// 	z01.PrintRune(char)
	// }
	// z01.PrintRune('\n')
}

func Wmatch(s1, s2 string) (res string) {
	for _, letter := range s1 {
		for _, char := range s2 {
			if char == letter {
				res += string(char)
				break
			}
		}
	}
	return res
}
