package main

import (
	"fmt"
)

func main() {
	fmt.Println(ActiveBits(7))
}

func ActiveBits(n int) int {
	var count int

	for n != 0 {
		count += n & 1
		n >>= 1
	}
	return count
}
