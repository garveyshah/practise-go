package main

import "fmt"
// Go supports methods on user-defined types
// The methods for a type are defined at the package block level

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

// Method declarations look like functions declarations, with one addition: the receiver specification
// They is one key difference between functions and methods:
 // Methods can only be defined ant the package block level, while functions can be  defined inside any block

 // Jusr like functions method names cannot be overloade. 
 // You can use the same method names for different types, but you can't use the same method
 // name for two different methods on the type 