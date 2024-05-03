// Write a Go function HasPrefix that takes a string and a prefix
// as input and returns true if the string starts with the given prefix,
// otherwise returns false.
package main

import "fmt"

func main() {
	prefix := "un"
	inputStr := "university"

	fmt.Println(HasPrefix(inputStr, prefix))
}

func HasPrefix(str, prefix string) bool {
	if len(str) < len(prefix) {
		return false
	}
	for i := 0; i < len(prefix); i++ {
		if prefix[i] != str[i] {
			return false
		}
	}
	return true
}
