package main

import "fmt"

func main() {
	fmt.Println(AddFront("Hello", []string{"world"}))
	fmt.Println(AddFront("Hello", []string{"world", "!"}))
	fmt.Println(AddFront("Hello", []string{}))
}

func AddFront(s string, slice []string) []string {
	if s == "" {
		return slice
	}

	slice = append([]string{s}, slice...)

	return slice
}
