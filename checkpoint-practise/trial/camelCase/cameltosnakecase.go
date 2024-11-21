package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := []string{
		"CamelCase",
		"camelCase",
		"HelloWorld",
		"132",
		" ",
		"",
		"A",
		"abcs",
		"AbC",
		"AbCEf",
		"abcAree",
		"ahe1Abde",
		"tesTing1",
		"SOME_VARIABLE",
		"ASuperLonGVariableName",
		"thisIsaTestOfCamelCase",
		"aA",
	}
	for _, arg := range args {
		challenge.Function("CamelToSnakeCase",CamelToSnakeCase, solutions.CamelToSnakeCase, arg)
	}
}

// func main() {
// 	fmt.Println(CamelToSnakeCase(" "))
// 	fmt.Println(CamelToSnakeCase(""))
// 	fmt.Println(CamelToSnakeCase("HelloWorld"))
// 	fmt.Println(CamelToSnakeCase("helloWorld"))
// 	fmt.Println(CamelToSnakeCase("camelCase"))
// 	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
// 	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
// 	fmt.Println(CamelToSnakeCase("hey2"))
// }

func CamelToSnakeCase(input string) string {
	if input == "" {
		return ""
	}

	if !ValidCase(input) {
		return input 
	}

	res := ""

	for i := 0; i < len(input); i++ {
		if i != 0 && Up(rune(input[i])) && i+1 < len(input) && Low(rune(input[i+1])) {
			res += "_" +  string(input[i])
		} else if  Low(rune(input[i])) || i == 0 && Up(rune(input[i])) {
			res += string(input[i])
		} else {
			return input
		}
	}
	return res
}

func ValidCase(str string) bool{
	for i, char := range str {
		if i != 0 && (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			return true
		} 
	}
	return false
}

func Up(char rune) bool {
	return char >= 'A' && char <= 'Z' 
}

func Low(char rune) bool {
	return char >= 'a' && char <= 'z' 
}