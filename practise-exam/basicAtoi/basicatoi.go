package main

import "fmt"

func BasicAtoi(s string) int {
	result := 0

	for _, char := range s {
		result = result*10 + int(char-'0')
	}
	return result

}
func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
}
