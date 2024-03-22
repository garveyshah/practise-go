package main

import "fmt"

func ToUpper(str string) string {
	result := ""
	for _, char := range str {
		if char >= 'a' && char <= 'z' {
			result += string(char - 'a' + 'A')
		} else {
			result += string(char)
		}
	}
	return result
}

func main() {
	str := "Hellow World!"
	fmt.Println(ToUpper(str))
}
