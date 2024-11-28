// Goroutine is an idependent function that excutes simultanuosly in some separete lightweight threads managed by go.

package main

import (
	"fmt"
	"time"
)

func main() {
	go helloworld()
	// time.Sleep(1 * time.Second)
	go goodbye()
	time.Sleep( 32 * time.Microsecond)
}

func helloworld() {
	fmt.Println("Hellow World")
}

func goodbye() {
	fmt.Println("Good Bye")
}
