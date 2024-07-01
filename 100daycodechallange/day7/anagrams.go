/*
Day 7: Implement a function to check if two strings are anagrams.
*/

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

	str1, str2 := GetInput(MainStr)
	// fmt.Printf("String 1 is :%v.\nString 2 is :%v.",Str1, Str2)

	if ConfirmAnagrams(str1, str2) {
		fmt.Printf("\n%v and %v are anagrams.\n", str1, str2)
	} else {
		fmt.Printf("\n%v and %v are NOT anagrams of each other.\n", str1, str2)
	}
}

// Function to format the input into two strings.
func GetInput(s string) (string, string) {
	strSlice := strings.Split(s, " ")

	str1, str2 := strings.Trim(strSlice[0], ""), strings.Trim(strSlice[1], "")
	return str1, str2
}

// Function that checks if two strings are anagrams.
func ConfirmAnagrams(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	countChar := make(map[rune]int)

	for _, char := range str1 {
		countChar[char]++
	}

	// Decrease count for each character in str2
	for _, char := range str2 {
		if countChar[char] == 0 {
			return false
		}
		countChar[char]--
	}

	// Check if all count are zero
	for _, count := range countChar {
		if count != 0 {
			return false
		}
	}
	return true
}
