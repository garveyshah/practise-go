package main

import (
	"fmt"
)

func main() {
	str1 := "abcdedfeddffnnf"
	str2 := "abfdsnngdf"
	fmt.Println(WeAreUnique(str1, str2))
}

func WeAreUnique(str1, str2 string) int {
	var duplicate bool = false
	var unique []rune
	var bros []rune

	for _, char := range str1 {
		for _, dup := range str2 {
			if char == dup {
				duplicate = true
			}
			if !duplicate {
				unique = append(unique, char)
			} else {
				bros = append(bros, dup)
			}
		}
	}
	fmt.Println(unique)
	fmt.Println(string(bros))
	return len(unique)
}
