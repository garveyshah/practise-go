/*
	Using 'reflect.Type' We can explore any kind of struct, with any number of fields and types.
	The example below uses a recursive type inspector that prints the structure of any given type.
	The inspector iterates through the struct fields even if they are other struct.
	This is done obtaining the `kind` of field and comparing it with a struct (f.Type.Kind() == reflect.Struct).
	Note that the code does not skip unexported fields.
*/

package main

import (
	"fmt"
	"reflect"
	"strings"
)

type T struct {
	B int
	C string
}

type S struct {
	C string
	D T
	E map[string]int
}

func printerReflect(offset int, typeOfX reflect.Type) {
	indent := strings.Repeat(" ", offset)
	fmt.Printf("%s %s: %s {\n", indent, typeOfX.Name(), typeOfX.Kind())
	if typeOfX.Kind() == reflect.Struct {
		for i := 0; i < typeOfX.NumField(); i++ {
			innerIndent := strings.Repeat(" ", offset+4)

			f := typeOfX.Field(i)
			if f.Type.Kind() == reflect.Struct {
				printerReflect(offset+4, f.Type)
			} else {
				fmt.Printf("%s %s: %s\n", innerIndent, f.Name, f.Type)
			}
		}
	}
	fmt.Printf("%s }\n", indent)
}

func main() {
	x := S{"root", T{42, "forty two"}, make(map[string]int),
}

printerReflect(0, reflect.TypeOf(x))
}