package main

import (
	"fmt"
	project01 "project01"
)

func main() {
	a := 0
	b := 1
	project01.Swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
