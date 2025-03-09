package main

import "fmt"

// An array - Is an indexed sequence of elements witha given length.
// Arrays are typed and their size is fixed.

func main() {
	var a[5] int
	fmt.Println(a)

	b := [5]int{0, 1,2,3,4}
	fmt.Println(b)

	c := [5]int{0,1,2}
	fmt.Println(c)
}

