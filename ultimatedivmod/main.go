package main

import (
	"fmt"
	"project"
)

func main() {
	a := 13
	b := 2
	project01.UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
