/* 
* Go permits the definition of anonymous structs like the shown in line 23.
Compared with a regular struct like `Circle` printing the struct brings a similar result.
However, we cannot print its name as we do with type `Circle`. The fields from the anonymouss
function can be modified like done with regular structs.
Anonymous structures can be compared with other structures if only they have the same fields
*/

package main

import (
	"fmt"
	"reflect"
)

type Circle struct {
	x      int
	y      int
	radius int
}

func main() {
	ac := struct {x int; y int;	radius int}{1, 2, 3}

	c := Circle{10, 10, 3}

	fmt.Printf("%+v\n", ac)
	fmt.Println(reflect.TypeOf(ac))

	fmt.Printf("c -> %+v\n", c) // Prints fields 
	fmt.Println(reflect.TypeOf(c)) // prints the name 

	ac.x=3
	fmt.Printf("ac.x %+v\n", ac)

	ac = c
	fmt.Printf("%+v\n", ac)

	fmt.Println(reflect.TypeOf(ac))
}
