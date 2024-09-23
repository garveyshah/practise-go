package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	for _, char := range args {
		fmt.Println(Reversestrcap(char))
	}
}

func Reversestrcap(str string) string {
	res := ""
	for i, char := range str {
		if i != len(str)-1 {

			if Up(char) && str[i+1] != ' ' {

				res += string(char + 32)

			} else if Low(char) && str[i+1] == ' ' {

				res += string(char - 32)

			} else {
				res += string(char)
			}
		} else {

			if i == len(str)-1 && Low(char) {
				res += string(char - 32)
			} else {
				res += string(char)
			}
		}
	}
	return res
}

func Low(s rune) bool {
	return s >= 'a' && s <= 'z'
}

func Up(s rune) bool {
	return s >= 'A' && s <= 'Z'
}
