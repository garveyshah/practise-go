package main

import (
	t "fmt"
)

// https://www.golangprograms.com/how-to-initialize-an-array-with-an-array-literal-in-go.html

/* YOU can initialize an array with pre-defined values using an array literal. An array literal have the number of elements it will hold in square brackets, followed by the type of its elements.
This is followed by a list of initial values seprated by commas of each element inside the curly braces.
*/

func main() {
	x := [5]int{10, 20, 30, 40, 50}   // Initialized with values
	var y [5]int = [5]int{10, 20, 30} // Partial assignment
	t.Println(x)
	t.Println(y)

}
