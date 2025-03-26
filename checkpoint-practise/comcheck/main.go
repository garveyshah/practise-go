package main

import (
	"os"
	t"github.com/01-edu/z01"
)

func main() {
	myWords := make(map[string]bool)
	words := [3]string{"01", "galaxy", "galaxy 01"}

	for _, word := range words {
		myWords[word] = true
	}

	input := os.Args[1:]

	for _, word := range input {
		if _, ok := myWords[word]; ok {
			Println("Alert!!!")
		}
	}
}

func Println(s string) {
	for _, char := range s {
		t.PrintRune(char)
	}
	t.PrintRune(10)
}