// 3.2.1 Length and capacity
// While an array allocates a fixed amount of memory that is directly related to its length.
// a slice can reserve a larger amount of memory that does not necceraily have to be filled.
package main

import "fmt"

// Prints the length and capacity of various variables 

func main() {
	a := []int{0, 1, 2, 3, 4 }

	fmt.Println("a >> ", a, len(a), cap(a))
	// Variable 'a' is an array and therefore, its length and capacity are the same.

	b := append(a,5)
	fmt.Println("b >> ", b, len(b), cap(b))
	// 'b' a slice built by appending number '5' to 'a' has a different length and capacity.


	c := b[1:4]
	fmt.Println("c >> ", c, len(c), cap(c))
	// Selecting elements of slice 'b' starting at index 1 will result in a slice with the original capacity minus one.

	d := make([]int, 5, 10)
	fmt.Println("d >> ", d, len(d), cap(d))
	// 'd' is built using the make function.The function takes a slice type, its length
	// and capacity as arguments

	// d[6]=5 // -> This will fail
	// Any element with a position equal to or greater than length cannot be accessed independently of the slice capacity.

}