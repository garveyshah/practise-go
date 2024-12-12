package main

type Person struct {
	FirstName string
	LastName  string
	Age       string
}

// This should be read as declaring a user-defined type with the name Person to have the Underlying type of the struct that follows.
// You can use any primitive type or compound type literal to define a concrete type.
type (
	Score     int
	Converter func(string) Score
	TeamScore map[string]Score
)

// Note
//  An 'Abstract type' is one that specifies what a type should do but not how it is done.
// A 'concrete type' specifies what and how. This means that the type has a specified way to store it's data and provides an implementation of any methods declared on the type.
// While all types in Go  are either abstract or concrete , some languages allow hybrid types,  such as abstract classes or interfaces with default methods in Java.
