package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	args := os.Args[1]

	// fmt.Println(args)

	word := ""
	WordsArr := []string{}

	for _, char := range args {
		if char == ' ' {
			if word != "" {
				WordsArr = append(WordsArr, word)
				word = ""
			}
		} else {
			word += string(char)
		}
	}

	if word != "" {
		WordsArr = append(WordsArr, word)
	} 
	
	fmt.Println(string(WordsArr[len(WordsArr)-1]))
}
