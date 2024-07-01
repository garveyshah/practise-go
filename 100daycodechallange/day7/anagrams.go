/*
Day 7: Implement a function to check if two strings are anagrams.
*/

// workflow
// 1. interate through the strings look for instances of "" if found append to separate string.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter the strings each enclosed in \"\" and separated by a space.")

	reader := bufio.NewReader(os.Stdin)
	MainStr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input", err)
		return
	}

	MainStr = strings.TrimSpace(MainStr)
	fmt.Println(MainStr)

}

func GetInput(s string) string{
	
	return ""
}
