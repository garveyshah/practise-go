/*
Frequency Counter - Write a function that takes a string and returns a map with the frequency of each character.
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Enter a string enclosed in quotation marks.")
		return
	}

	fmt.Println(FrequencyCounter(os.Args[1]))
}

func FrequencyCounter(s string) map[string]int {
	WordLib := make(map[string]int)

	for _, char := range s {
		WordLib[string(char)]++
	}
	return WordLib
}
