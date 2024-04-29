package main

import "fmt"

func Concat(s, y string) string {
	concatR := string(s) + string(y)
	return concatR
}

func main() {
	y := "World!"
	s := "Hello "

	fmt.Println(s, y)
}
