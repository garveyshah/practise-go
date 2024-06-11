package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1]

	str := ""
	for i := len(args) - 1; i >= 0; i-- {
		if string(args[i]) == " " && str == "" {
			continue
		}
		if string(args[i]) != " " {
			str = string(args[i]) + str
			// fmt.Println(str)
		}
		if string(args[i]) == " " {
			break
		}

	}
	if str == "" {
		//str = "usage: type a string"
		return
	}

	for _, ch := range str {
		z01.PrintRune(ch)
	}
	z01.PrintRune(10)
}
