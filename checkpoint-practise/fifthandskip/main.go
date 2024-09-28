package main

import (
	"fmt"
)

func main() {
	fmt.Println(FifthAndSkip("e 5£ @ 8* 7 =56 ;"))
}

func FifthAndSkip(str string) string {
	if str == "" {
		return "\n"
	}

	if len(str) < 5 {
		return "Invalid Input\n"
	}

	str1 := Clean(str)

	clean := ""
	count := 0

	for _, char := range str1 {
		if count == 5 {
			clean += string(' ')
			count = 0
		} else {
			clean += string(char)
			count++

		}
	}
	return clean + "\n"
}

func Clean(s string) (res string) {
	for _, char := range s {
		if char == ' ' {
			continue
		} else {
			res += string(char)
		}
	}
	// fmt.Println(res)
	return res
}
