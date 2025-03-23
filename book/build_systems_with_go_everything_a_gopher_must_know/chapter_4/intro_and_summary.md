# CHAPTER 4
## STRUCTS, METHODS, AND INTERFACES

## STRUCTS
A `struct` is a sequence of elements named fields.
Each field has a name and a type.

- Not passing any argument initializes every field in a `struct` with the zero value.
- Intial values can be set passing them as arguments in the same order they were declared.
- We can set what value corresponds to what field using the field name.
- In Go it does not exist the concept of a constructor like in other languages.
- A `struct` is a very flexible construct that can be defined in many ways.
- Go permits the definition of anonymous structs.
- Anonymous structs can be compared with other structs if and only if they have the same fields.
- Structs can be nested to incorporate other structs definition.
- To embed a struct in other structs, this has to be declared as a nameless field.

## METHODS
Go defines methods as a special function with a receiver. 
The receiver sets the type that can invoke that function.
- Recivers define the type `receiving` the logic inside a method.
- Embedding structs can be done only if the compiler find no ambiguities.
- When a struct is embedded into other structs its methods are made available to the second one.

## INTERFACES
An `interface` is a collection of methods with any signature.
Interfaces do not define any logic or value, they simply define a collection of methods to be implemented.

- `Empty Interface` this interface has no method and is implemented by any type.
-The empty interface is a very ambigouse context for a typed language like Go.

Go cannot be fully considered an object oriented language. Actauly concepts such as classes or hierarchy of classes do not exist. Similar functionality can indeed be obtained using the current language definiation. 
However, we cannot say Go is an object-oriented language

# SUMMARY
This Chapter introduces the concepts of struct, methods, and interfaces used in Go. These concepts are fundamental pieces of the Go language. Additionaly, certain situations that may seem weird to early adopters such as the difference between value and pointer receivers in methods are explained.
