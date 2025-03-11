/*
	To embed a struct in other structs, this has to be decalred  as a nameless field.
	In the example below, by embedding the `Coordinates` struct in the `Circle` type we make
	fiels `x` and `y` directly accessible.
	Thr coordinates instance can also be accessed like the `Coordinates` field.
*/

package main

import "fmt"

type Coordinates struct {
	x int
	y int
}

type Circle struct {
	Coordinates
	Radius int
}

type A struct { fieldA int }

type B struct { fieldA int }

type C struct {
	A
	B
}


func main() {
	c := Circle{Coordinates{1, 2}, 3}
	fmt.Printf("%+v\n", c)

	fmt.Printf("%+v\n", c.Coordinates)
	fmt.Println(c.x, c.y)


	a := A{10}
	b := B{20}
	e := C{a, b}
	 
	// Ambigous access
	// fmt.Println(e.fieldA) - ambigous selector
	fmt.Println(e.A.fieldA, e.B.fieldA)
}
