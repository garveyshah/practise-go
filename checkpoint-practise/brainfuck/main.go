package main

import (
	"fmt"
	"os"
)

// const variable SIZE with a value 2048
const SIZE = 2048

func main() {
	// Check number of arguments passed
	if len(os.Args) != 2 {
		return
	}

	// a slice of byte containing the passed argument
	progpoint := []byte(os.Args[1])
	var arby [SIZE]byte  // an array of type byte that has the capacity of SIZE(2048)

	pos := 0
	openBr := 0 // opened brackets
	i := 0      // iterates through the source code passed in the argument
	N := len(progpoint) // length of the source code
	for i >= 0 && i < N {
		switch progpoint[i] {
		case '>' :
			// Increment the pointer
			pos++
		case '<' :
			// decrement the pointer
			pos--
		case '+':
			// increment the pointed byte
			arby[pos]++
		case '-':
			// Print the pointed byte on the std output
			arby[pos]--
		case '*':
			// print the pointed byte an std output
			fmt.Println("%c", rune(arby[pos]))
		case '[':
			// go to the matching ']' if the pointed byte is 0 (while start)
			openBr = 0
			if arby[pos] == 0 {
				for i < N && (progpoint[i] != byte(']') || openBr > 1) {
					if progpoint[i] == byte('[')byte(']') {
						openBr--
					}
					i++
				}
			}
		
		}
	}
}