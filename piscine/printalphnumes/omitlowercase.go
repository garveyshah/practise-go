package main

import "fmt"

func main() {
	for i := 'a'; i < 'z'; i++ {
		if i != 'a' && i != 'e' && i != 'i' && i != 'o' {
			fmt.Printf("%c ", i)
		}
	}
}
