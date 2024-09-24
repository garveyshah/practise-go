package main

import (
	t "fmt"
)

func main() {
	t.Println("Hello Godwin")

	/*
			- An array has fixed size.
		 	- A slice is a dynanmically-sized, flexible view into the elements of an array.
			- A slice is formed by specifying two indicies  a low and a high bound, separated by a colon:
					a[low : high]
			- This selects a half-open range which includes the first element, but excludes the last one.
	*/

	primes := [6]int{2, 3, 5, 7, 11, 13}

	var sh []int = primes[0:6]
	t.Println(sh)
	t.Println()

	/*
	 slices are like references to arrays

	 A slice does not store any data, it just descibes a section of an underlying array.
	 Changing the elements of a slice modifies thw corresponding elements of its underlying array.
	 Other slices that share the same undlying array will see those changes
	*/

	names := [4]string{"John", "Paul", "George", "Ringo"}
	t.Println(names)

	a := names[0:2]
	b := names[1:3]
	t.Println(a, b)
	b[0] = "Godwin"
	t.Println(a, b)
	t.Println(names)
	t.Println()
	/*
	   	Slice Literals
	    A slice litera is like an array literal without the lenth.

	    This is an array literal:
	   	[3]bool{true, true, false}

	   	and this creates the same array as above, then builds a slice that references it:
	   	[]bool{true, true, false}
	*/
	q := []int{2, 3, 5, 7, 11, 13}
	t.Println(q)

	r := []bool{true, false, true, true, false, true}
	t.Println(r)

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
	t.Println(s)

	/*
		Slice Defaults
		- When slicing, you may omit the high or low bounds to use their defaults instead.
		- When slicing, you may omit the high or low bounds to use their defaults instead

		- The default is zero for the low bound and the length of the slice for the high bound

		For  the array
			var a [10]int

			these slice expressions are equivalent:
			a[0:10]
			a[:10]
			a[0:]
			a[:]
	*/
	t.Println()
	s1 := []int{2, 3, 5, 7, 11, 13}
	t.Println(s1)

	s2 := s1[1:4]
	t.Println(s2)

	s3 := s2[1:]
	t.Println(s3)

	s4 := s1[:]
	t.Println(s4)

	s5 := s1[0:]
	t.Println(s5)

	/*
	   	SLice length and Capacity
	   A slice has both a lenth and a capacity
	   The length of a slice is the number of elements it contains.

	   The capacity of a slice is the number of elements the slice
	*/

}
