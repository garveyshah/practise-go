// Channels
// In concurrent programming, Go provides channels that you can use for biredirectional communication btwn goroutines.

// Bidirectional communication means that one goroutine will send a message and the other will read it.
// Sends and receives are blocking. Code execution will be stopped until the write an read done successfully

// Channels are one of the more convinient ways to send and receive notifications.

// Types of Channels
// 1. Unbuffered channel 
//  Unbeffered channels require both the sender and receiver to be present to be successful operations.
// It requires a goroutine to read the data, otherwise, it will lead to deadlock. 
// 	Channels are unbuffered by default

// 2. Buffered channels:
//  Buffered channels have the capacity to store values for future processing. 
// The sender is not blocked until it becomes full and it doesn't necessarily need a 
// reader to complete the synchronization with every operation.

// If a space in the array is available, the sender can send its value to the channel and complete 
// its operation immediately.



package main

import (
	"fmt"
	"time"
)

func main() {
	
	msg := make(chan string) // initialize a channel

	go greet(msg) // start a go routine

	time.Sleep(2 * time.Second) // delay execution of the main by 2 second 

	greeting := <- msg // store

	// time.Sleep(2 * time.Second)  
	fmt.Println("Greeting received")
	fmt.Println(greeting)
}

func greet(ch chan string) {

fmt.Println("Greeter waiting to send greeting")

ch <- "Hello Rwitesh"

fmt.Println("Greeter completed")
}

// msg := make(chan string) is decalaring a channel of type string.
// Then the channel is passed in goroutine greet. 
// 'ch' '<- "Hello Rwitesh" allows us to write the message to 'ch'