/*
	Using 'reflect.value' we can set values on run-time.
	Every 'Value' has methods 'setInt32, setFloat32, setting, etc.
	That set the field to a 'int32', 'float32', 'string', etc. value
	the example below sets the string fields from a struct to uppercase
*/
package main

import (
	"fmt"
	"reflect"
	"strings"
)

type T struct {
	A string 
	B int
	C string
}

func main() {
	t := T{"Hello", 42, "bye"}
	fmt.Println(t)
	
	valueOfT := reflect.ValueOf(&t).Elem()
	for i := 0; i < valueOfT.NumField(); i++ {
		f := valueOfT.Field(i)
		if f.Kind() == reflect.String {
			current := f.String()

			f.SetString(strings.ToUpper(current))
		}
	}
	fmt.Println(t)

}