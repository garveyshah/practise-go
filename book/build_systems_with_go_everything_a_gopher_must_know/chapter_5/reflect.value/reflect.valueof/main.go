/*
	We have seen how 'reflect.Type' permits code introspection.
	However, the 'reflect.Type' type cannot acces field values.
	This functionality is reserved to the 'reflect.Value' type.
	Actualyy, from a 'reflct.Value' type we can access its 'reflect.Type' type.

	
*/

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int32 = 42
	var b string = "forty two"

	valueOfA := reflect.ValueOf(a)
	fmt.Println(valueOfA.Interface())

	valueOfB := reflect.ValueOf(b)
	fmt.Println(valueOfB.Interface())
}


