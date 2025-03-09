// 3.2.1 Length and capacity
// The common operation on a slice is the iteration through its items.
// Any 'loop' is a good candidate to do this
// For arrays and slices, range returns the index of the item and the item itself for every iteration of the loop.
// The iteration is always done incrementally from index 0 to index n-1

package main

import "fmt"

func main() {
	names := []string{
		"Jeremy", "John", "Joseph",
	}

	// for i := 0; i < len(names); i++ {
	// 	// fmt.Println(i, names[i])
	// 	name = name + "_changed"
	// }

	for position, name := range names {
		// fmt.Println(position, name)
		names[position] = name + "_changed"
	}
	fmt.Println(names)
}