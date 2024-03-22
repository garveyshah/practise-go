package main

import "fmt"

func CapitalizeEachWord(s string) string {
	result := ""
	capitalizeNext := true

	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			if capitalizeNext {
				if char >= 'a' && char <= 'z' {
					result += string(char - 32)
					capitalizeNext = false
				}
			} else {
				result += string(char)
			}
		} else {
			result += string(char)
			capitalizeNext = true
		}
	}
	return result
}

func main() {
	string := "hello world! my name is godwin. ouma but, user name is oumouma."
	fmt.Println(CapitalizeEachWord(string))
}
