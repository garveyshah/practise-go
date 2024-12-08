// Closing the channel:

// CLosing the channel indicates that no more values should be sent on it.
// We want show the work has been completed and there is no need to keep a channel open.

package main

import (
	"fmt"
	"time"
)

func main() {
	msg := make(chan string)
	go greet(msg)

	time.Sleep(2 * time.Second)

	greeting := <- msg
	fmt.Println("Greeting received")
	fmt.Println(greeting)

	// _, ok := <- msg
	// if ok {
	// 	fmt.Println("Channel is open!")
	// } else {
	// 	fmt.Println("Channel is closed!")
	// }

	msg1 := <- msg
	fmt.Println(msg1)


	msg2 := <- msg



	fmt.Println(msg2)
}

func greet(ch chan string) {
	time.Sleep(4 * time.Second)
	fmt.Println("Greeter waiting to send greeting!")

	ch <- "Hello Rwitesh"
	ch <- "I love my job"
	// close(ch)

	fmt.Println("Greeter completed")
}
