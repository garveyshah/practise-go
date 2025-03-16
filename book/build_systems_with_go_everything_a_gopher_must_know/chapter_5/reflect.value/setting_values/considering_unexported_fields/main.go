/*
	If the set operation is not valid, the operation will panic.
	For example setting fields to a differemt type or trying to set unexported fields.
	In the example below, the field 'c' is unexported. 
	Additional checking must be done using the 'canSet()' method.
	Using this method we can skip unexported fields.
	Observe that the output has not modified 'c' value
*/

package main

import (
	"fmt"
	"reflect"
	"strings"
)

type T struct {
	A string
	B int32
	c string
}

func main() {
	t := T{"hello", 42, "bye"}
	fmt.Println(t)

	valueOfT := reflect.ValueOf(&t).Elem()
	for i := 0; i < valueOfT.NumField(); i++ {

		f := valueOfT.Field(i)

		if f.Kind() == reflect.String {
			if f.CanSet() {
				current := f.String()

				f.SetString(strings.ToUpper(current))
			}
		}
	}
	fmt.Println(t)
}