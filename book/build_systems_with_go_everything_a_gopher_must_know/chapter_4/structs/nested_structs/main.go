/*
	Structs can be nested to incorporate other structs definition.
	The example below defines a `Circle` type using a `Coordinates` type.
*/

package main

import "fmt"

type Coordinates struct {
	x int
	y int
}

type Circle struct {
	center Coordinates
	radius int
}

func main() {
	c := Circle{Coordinates{1, 2}, 3}
	fmt.Printf("%+v\n", c)
}
