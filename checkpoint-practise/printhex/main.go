package main

import (
	"os"

	"project01/checkpoint-practise/atoi/atoi"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	num, err := atoi.Atoi(os.Args[1])
	if err != nil {
		Print("ERROR")
		return
	}

	hex := DigitToHex(num)
	Print(hex)
}

func DigitToHex(num int) string {
	return ""
}

func Print(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
	z01.PrintRune(10)
}
