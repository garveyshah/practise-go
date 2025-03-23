/* 
	The example below uses `reflect.TypeOf` to access all the declared fields of the struct `user`. 
	The type `Type` returned by `TypeOf` has functions to check the type name, number of fields, size, etc.
	Field information is stored in a type `StructField` that can be accessed using `Field()` and `FieldByName()` functions.
	For every `StructField` tags are stored in a `StructTag` type ('fieldUserId.Tag')
	A 'StructTag' contains all the available tags of a field.
*/
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	UserId string `tagA:"valueA1" tagB:"valueA2"`
	Email string `tagB:"value"`
	Password string `tagC:"v1 v2"` 
}

func main() {
	T := reflect.TypeOf(User{})

	fieldUserId := T.Field(0)
	t := fieldUserId.Tag
	fmt.Println("StrucTag is :", t)
	v, _
}
