/*
	Beyond type exploration, we can check if a type implements an interace.
	This can be done using `Implements` method as shown below.
	This is good example of how interfaces works in Go.
	The method `Add` has a pointer receiver (*Calculator) for that reason, the `main.Calculator` type does not implement the `Adder` interface.
*/

package main

import (
	"fmt"
	"reflect"
)

type Adder interface {
	Add(int, int) int
}

type Calculator struct{}

func (c *Calculator) Add(a int, b int) int {
	return a + b
}

func main() {
	var ptrAdder *Adder
	adderType := reflect.TypeOf(ptrAdder).Elem()

	c := Calculator{}

	calcType := reflect.TypeOf(c)

	calcTypePtr := reflect.TypeOf(&c)

	fmt.Println(calcType, calcType.Implements(adderType))
	fmt.Println(calcTypePtr, calcTypePtr.Implements(adderType))
}
