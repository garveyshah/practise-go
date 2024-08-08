package main

import (
	"fmt"
)

func main() {
	elems := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(BasicJoin(elems))
}

func BasicJoin(elems []string) string {
	var result string
	for _, str := range elems {
		result += str
	}
	return result
}
