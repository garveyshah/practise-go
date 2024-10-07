package main

import (
	"fmt"
	"os"
	"project01/100daycodechallange/day18/checker"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage : go run . <Words seperated by spaces>\n\nExample : go run . preprocess precompute prefix preview")
		return
	}
	words := os.Args[1:]

	fmt.Printf("The Longest prefix is %q\n",checker.Prefix(words))
}
