package main

import "fmt"

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}

func HashCode(dec string) string {
	var result string

	for _, char := range dec {
		if char > 126 || char < 32 {
			result += string((char + rune(33)+ rune(len(dec)))%127)
		} else {
			result += string((char + rune(len(dec))) % 127)
		}
	}
	return result
}
