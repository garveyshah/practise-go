/*

https://www.programiz.com/golang/goroutines 

In Go, goroutines are used to create concurrent programs(programs that are able to run multiple processes at the same time).
We can convert a regular function to a goroutine by calling the function using the 'go' keyword.
For Example:

	func display() {
		// code inside function
	}

// start goroutine
go dispaly()

Here, display() is a goroutine, belkow is a working example.
*/
package main

import (
	"fmt"
)

// / create a function
func dispaly(message string) {
	fmt.Println(message)
}

func main() {
	// call goroutine
	go dispaly("Process 1")
	dispaly("process 2")
}

// Output: 
// Process 2
