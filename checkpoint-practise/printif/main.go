package main

import (
	"fmt"
)

func main() {
	fmt.Print(PrintIf("abcdefz"))
	fmt.Print(PrintIf("abc"))
	fmt.Print(PrintIf(""))
	fmt.Print(PrintIf("14"))
}

func PrintIf(s string) string {
	var result string
	if len(s) >= 3 {
		result = "G\n"
	} else if s == "" {
		result = "G\n"
	} else {
		result = "Invalid Input\n"
	}
	return result
}