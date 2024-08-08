package main

import (
	"fmt"
)

func main() {
	fmt.Println(Sqrt(3))
	fmt.Println(Sqrt(100000000))
}

func Sqrt(nb int) int {
	if nb == 0 {
		return 0
	}

	var result int = nb

	for i := 0; i < 1000; i++ {
		result = (result + nb/result) / 2
	}
	if result*result != nb {
		return 0
	}
	return result
}
