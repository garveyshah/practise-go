/*
    In the object-oriented world, a method is defined as a procedure associated with a class.
Go doesn't have such a thing as classes.
However, Go defines methods as a special function with a receiver.
The reciver sets the type that can invoke that function.

Assume we work with the `Rectangle` type and we want to add some operations such as computing the surface.
In the example in `definition_of_methods` `Surface() int` is a function with the receiver `(r Rectangle)`. 
This means that any type `Rectangle` van invoke this method.
Inside the method, the fields `Height` and `Width` from the current instance `R` are accessible.
*/

