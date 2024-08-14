/*
Anagram Checker - Write a function to check if two strings are anagrams of each other.
*/

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		Print("Enter the two strings to check if they are anagrams of each other. Separated by a space")
		return
	}

	str1 := os.Args[1]
	str2 := os.Args[2]

	if AnagramChecker(str1, str2) {
		Print(str1 + " and " + str2 + " are anagrams of each other.")
	} else {
		Print(str1 + " and " + str2 + " are NOT anagrams of each other.")
	}
}

func Print(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
	z01.PrintRune(10)
}

// // Function AnagramChecker checks if two strings are Anagrams of each other or not.
// func AnagramChecker(a, b string) bool {
//	// If the lengths are not equal, they can't be anagrams
// 	if len(a) != len(b) {
// 		return false
// 	}

// 	letters := make(map[rune]int)

// 	// Count each letter in string a
// 	for _, char := range a {
// 		letters[char]++
// 	}

// 	// Subtract each letter in string b
// 	for _, char := range b {
// 		letters[char]--
// 		if letters[char] < 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// Function AnagramChecker checks if two strings are Anagrams of each other or not.
func AnagramChecker(a, b string) bool {
	// If the lengths are not equal, they can't be anagrams
	if len(a) != len(b) {
		return false
	}

	letters := make(map[rune]int)

	// Count each letter in string a
	for _, char := range a {
		letters[char]++
	}

	// Subtract each letter in string b
	for _, char := range b {
		// If char is not in the map or count goes below zero, return false
		if _, ok := letters[char]; !ok || letters[char] == 0 {
			return false
		}
		letters[char]--
	}
	return true
}
