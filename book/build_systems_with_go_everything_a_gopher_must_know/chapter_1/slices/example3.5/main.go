// 3.2.1 Length and capacity
// The common operation on a slice is the iteration through its items.
// Any 'loop' is a good candidate to do this

package main

import "fmt"

func main() {
	names := []string{
		"Jeremy", "John", "Joseph",
	}

	for i := 0; i < len(names); i++ {
		fmt.Println(i, names[i])
	}

	for position, name := range names {
		fmt.Println(position, name)
	}
}