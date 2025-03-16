/*
The example below uses 'reflect.Value' with a struct to print the field values.
The reflected value of variable 't' is correctly printed.
Similary, we can print the value of every field in the struct. Notice the difference between printing
'field.String()' and 'field'. For numeric values 'field.String()' returns a string like '<int32 Value>'
The string informs that there is an integer value in that field.
However, 'fmt.Println(field)' works as expected.
This occurs because the function prints the corresponding value when it receives 'value' types.
*/
package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A int32
	B string
	C float32
}

func main() {
	t := T{42, "forty two", 3.14}

	valueT := reflect.ValueOf(t)
	fmt.Println(valueT.Kind(), valueT)

	for i := 0; i < valueT.NumField(); i++ {
		field := valueT.Field(i)
		fmt.Println(field.Kind(), field.String(), field.Interface())
	}
}
