// Slice - defined as a "descriptor" for a contigous segment of an underlying array and provides access
//  to a numbered sequence elements from that array
// A slice is a refrence to an array.
// Does not store any dara but offers a view of it.

/* 
	A[0] Elements at position 0
	a[3:5] Elements from position 3 to 4
	a[3:] All the elemets starting at position 3
	a[:3] All the elements from  the start till position 2
	a[:]  All the elements

*/

package main

import "fmt"

func  main() {
	a := [5]string{"a","b","c","d","e"}

	fmt.Println("a => ", a)
	fmt.Println("a[:] => ",a[:])
	fmt.Println("a[0] => ",a[0])
	fmt.Println("a[0], a[1], a[2], a[3], a[4] => ",a[0], a[1],a[2], a[3], a[4])
	fmt.Println("a[0:2] => ",a[0:2])
	fmt.Println("a[1:4] => ",a[1:4])
	fmt.Println("a[:2] => ",a[:2])
	fmt.Println("a[2:] => ",a[2:])
}
