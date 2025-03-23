// A program to count the number of arguments 
package main

import (
	"fmt" // for output to the command-line
	"os" // for reading commandline arguments
)

func main() {
	// Check if number of arguments passed are more than 2 to pass
	if len(os.Args) < 2 {
		fmt.Println(0)
		return
	}

	
	count := 0 // initialize an int variable count 
	
	for _, arg := range os.Args[1:] {  // range of the slice of cmd arguments
		count += len(arg)  // Add the length of each arg to the variable count
	}
	
	fmt.Println(count) // Print the total count
}