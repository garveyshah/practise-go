// In Go most of the time, we work with 'slice' rather than arrays.
// Working with arrays or slice is very similar .
// Arrays are defined with a fixed length but slices are defined without a deffinate size.
package main

import (
	"fmt"

	"reflect"
)

func main() {
	a := [5]string{"a", "b", "c", "d", "e"}

	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(a[0:3]))
	fmt.Println(reflect.TypeOf(a[0]))
}