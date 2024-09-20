package main

import (
	"fmt"
	"project01/checkpoint-practise/canjump/sln1"
)

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(sln1.CanJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(sln1.CanJump(input2))

	input3 := []uint{0}
	fmt.Println(sln1.CanJump(input3))
}
