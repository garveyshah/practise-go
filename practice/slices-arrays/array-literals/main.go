package main

import (
	t "fmt"
)

// https://www.golangprograms.com/how-to-initialize-an-array-with-an-array-literal-in-go.html

/* YOU can initialize an array with pre-defined values using an array literal. An array literal have the number of elements it will hold in square brackets, followed by the type of its elements.
This is followed by a list of initial values seprated by commas of each element inside the curly braces.

1. Array Literals 
- An array literal specifies the elements of the array enclosed in curly braces {}
The length of the array is either specifeied explicitly or inferred from the number of elements.

Syntax
var arrayName = [lengh]Type{element1, element2, ...}

- You can either specify the length of the array or let Go infer it.
Key Points about Arrays:
- Arrays in Go have a fixed size, and the size is part of the array type.
- If the length is not specified explitly (like [...], it will be inferred from the number of elements provided.)

2. Slice Literals
- A slice is a dynamically sized, more flexible type built on top of arrays. Slice literals look similar to array literals but without specifying the size.
Syntax
	var sliceName = []Type{element1, element2, ...}

- In a slice literal, you omit the length, making it flexible.

Key Points about slices
 - Slices have a length that can be changed dynamically, unlike arrays.
 - You can append elements to slices, and the slice can grow in size.

*/

func main() {
	x := [5]int{10, 20, 30, 40, 50}   // Initialized with values
	var y [5]int = [5]int{10, 20, 30} // Partial assignment
	t.Println(x)
	t.Println(y)

var arr1 = [3]int{1,2,3} // Array of length 3 with specified elements
var arr2 = [...]string{"a", "b", "c"} // Array with inferred length
t.Println(arr1, arr2)

var slice1 = []int{1, 2, 3} 	// slice with 3 elements
var slice2 = []string{"Go", "Rust"} 	// Slice of strings
t.Println(slice1, slice2)

var numbers = []int{1,2,3}
numbers = append(numbers, 4, 5)
t.Println(numbers)
}
