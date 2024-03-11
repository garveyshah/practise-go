package main

import (
	"fmt"
	"project"
)

func main() {
	a := 13
	b := 2
	var div int
	var mod int
	project01.DivMod(a, b, &div, &mod)
	fmt.Println(div)
	fmt.Println(mod)
}
