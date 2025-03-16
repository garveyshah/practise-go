/*
	In order to know what type implements a value, we can compare iy with the 'Kind' type
	returned by method 'Kind()'.
	The type 'Kind' is a number representing one of the types available in Go (int43, string, map, etc).
	This cam be combined with a 'switch' statements as shown in Example below to identify what type are we working with.

*/

package main

import (
	"fmt"
	"reflect"
)

func ValuePrint(i interface{}) {
	v := reflect.ValueOf(i)

	switch v.Kind() {
	case reflect.Int32:
		fmt.Println("Int32 with value", v.Int())
	case reflect.String:
		fmt.Println("String with value", v.String())
	default:
		fmt.Println("Unkown type")
	}
}

func main() {
	var a int32 = 42
	var b string = "forty two"

	ValuePrint(a)
	ValuePrint(b)
}