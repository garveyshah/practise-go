/*
	The example below launches the `ShowIt` function in a goroutine that runs concurrenly with the `main` function. 
	Both functions print a message after sleeping (`time.Sleep`). 
	The `main` function sleeps half the time of `showIt` that is why we have a ratio of two messages from one function versus the other.
	It is important to notice that although the loop in `ShowIt` is endless, the program execution terminates when the `main` function finished.
	No goroutine will remain running when the 	`main` function finishes. 
*/

package main

import (
	"fmt"
	"time"
)

func ShowIt() {
	for {
		time.Sleep(time.Millisecond * 100)
		fmt.Println("Here it is!!!")
	}
}

func main() {
	go ShowIt()

	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 50)
		fmt.Println("Where is it?")
	}
}
