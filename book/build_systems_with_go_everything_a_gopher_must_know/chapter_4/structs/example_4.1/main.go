package main

import "fmt"

// Regresents the geometric figure with 'Height' and 'Width' fields
type Rectangle struct {
	Height int
	Width  int
}

func main() {
	a := Rectangle{} // Intializes every fieled in the struct with zero value
	fmt.Println(a)

	b := Rectangle{4, 4} //  Initial values can be set.
	fmt.Println(b)

	c := Rectangle{Width: 10, Height: 3} // setting what value corresponds to what field using the field name
	fmt.Println(c)

	d := Rectangle{Width: 7} // missing fields are set to there default value
	fmt.Println(d)
}