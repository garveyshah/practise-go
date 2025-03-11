/*
	The 'switch' Statement can extract the type of the value in runtime.
	This can be used to define how to operate with the value. In the example,
	we modify a print statement accordingly to the type.
	Notice that "%T" in the print statement gets the name of the value type.
*/
package main

import "fmt"

func main() {
	aux := []interface{}{42, "Hello", true}

	for _, i := range aux {
		switch t := i.(type) {
		default:
			fmt.Printf("%T -> %s\n", t, i)
		case int:
			fmt.Printf("%T -> %d\n", t, i)
		case string:
			fmt.Printf("%T -> %s\n", t, i)
		case bool:
			fmt.Printf("%T -> %v\n", t, i)
		}
	}
}
