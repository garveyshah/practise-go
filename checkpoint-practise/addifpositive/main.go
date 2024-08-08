package main

import "fmt"

func main() {
	fmt.Println(AddIfPositive(1, 2))
	fmt.Println(AddIfPositive(1, -2))
	fmt.Println(AddIfPositive(-1, 2))
	fmt.Println(AddIfPositive(-1, -2))
	fmt.Println(AddIfPositive(10, 20))
	fmt.Println(AddIfPositive(0, 20))
}

func AddIfPositive(a int, b int) int {
	var result int
	if a < 0 || b < 0 {
		return 0
	} else {
		result = a + b
	}
	return result
}
