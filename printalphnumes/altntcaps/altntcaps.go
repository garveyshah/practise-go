package main

import "fmt"

func main() {
	for i := 0; i < 26; i++ {
		if i%2 == 0 {
			fmt.Printf("%c", i+'a')
		} else {
			fmt.Printf("%c", i+'A')
		}
	}
	fmt.Println()
}
