/*
Palindrome Checker - Implement a function to check if a given string is a palindrome.
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}

	str1 := ToLower(os.Args[1])

	if PalindromeChecker(str1) {
		fmt.Printf("%v is a palindromes.\n", str1)
	} else {
		fmt.Printf("%v is NOT palindromes.\n", str1)
	}
}

func PalindromeChecker(a string) bool {
	new := ""
	for _, char := range a {
		if char >= 'a' && char <= 'z' {
			new += string(char)
		}
	}
	// fmt.Println(new)
	var checker bool
	for i := 0; i < len(new)/2; i++ {
		j := len(new) - 1 - i
		if new[i] != new[j] {
			checker = false
			break
		}
		checker = true
	}
	return checker
}

func ToLower(s string) string {
	result := ""
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			char += 32
		}
		result += string(char)
	}
	return result
}
