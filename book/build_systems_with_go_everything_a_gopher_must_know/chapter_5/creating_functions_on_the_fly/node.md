# CREATING FUNCTIONS ON THE FLY
In the other sections, we have explored how to inspect fields and modify values. Additionally, the `reflect` package permits the creation of new entities such as functions on the fly.
This offers certain functionalites available in other programming languages.
For example, the lack of generics in Go imposes some limitations although there is already a proposal at the moment of this writing.

Assume we want to write a function that generalizes the add operation. 
This function must sum numbers(integers and floats) and append strings. Given the current language defination, this is not posibble.
Check the code below. 
Go does not permit any kind of function overload.
Every function must have a unique name. Similarly, the lack of generics makes it impossible to reuse the same code using the add operator (+) defined for every type.

```go
func Sum(a int, b int) int {...}

func Sum(a float32, b float32) float32 {...} // Not unique.

func Sum(s string, b string) string {...} // Not unique
```
One Intresting walkarounf is to use the `reflect.MakeFunc` function to build our own functions with different signatures. 