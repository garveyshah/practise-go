package main

import "fmt"

func main() {
	fmt.Println("Hello Godwin")

	/*
			- An array has fixed size.
		 	- A slice is a dynanmically-sized, flexible view into the elements of an array.
			- A slice is formed by specifying two indicies  a low and a high bound, separated by a colon:
					a[low : high]
			- This selects a half-open range which includes the first element, but excludes the last one.
	*/

	primes := [6]int{2, 3, 5, 7, 11, 13}

	var sh []int = primes[0:6]
	fmt.Println(sh)
	fmt.Println()

	/*
	 slices are like references to arrays

	 A slice does not store any data, it just descibes a section of an underlying array.
	 Changing the elements of a slice modifies thw corresponding elements of its underlying array.
	 Other slices that share the same undlying array will see those changes
	*/

	names := [4]string{"John", "Paul", "George", "Ringo"}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)
	b[0] = "Godwin"
	fmt.Println(a, b)
	fmt.Println(names)
	fmt.Println()
	/*
	   	Slice Literals
	    A slice litera is like an array literal without the lenth.

	    This is an array literal:
	   	[3]bool{true, true, false}

	   	and this creates the same array as above, then builds a slice that references it:
	   	[]bool{true, true, false}
	*/
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
