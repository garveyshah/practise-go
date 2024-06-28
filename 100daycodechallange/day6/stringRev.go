/*
	Day 6: Write a program to reverse a string.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter the string to work on:")

	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error!! reading input", err)
		return
	}

	str = strings.TrimSpace(str)

	fmt.Println()
	fmt.Println(StrRev(str))
}
// function StrRev reverses a string.
func StrRev(s string) string {
	var newStr string
	for i := len(s) - 1; i >= 0; i-- {
		newStr += string(s[i])
	}
	return newStr
}
