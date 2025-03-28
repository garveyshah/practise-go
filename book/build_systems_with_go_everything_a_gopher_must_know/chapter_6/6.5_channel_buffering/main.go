/*
	In this example two functions and the `main` send data to the same channel.
	Due to the code workflow, the main writes to the channel when nobody is listening which triggers the error
	'fatal error: all goroutines are asleep - deadlock!.'
	In this case, a buffered channel can store messages until the receivers are ready to consume the messages.
	This example only needs a one-element buffer. However, you can check how removing the size value in the make 'make' statement returns an unbuffered channel
*/

package main

import (
	"fmt"
	"time"
)

func MrA(ch chan string) {
	time.Sleep(time.Millisecond * 500)

	ch <- "This is MrA"
}

func MrB(ch chan string) {
	time.Sleep(time.Millisecond * 200)

	ch <- "This is MrB"
}

func main() {
	ch := make(chan string)
	// ch := make(chan string, 1)

	go MrA(ch)
	go MrB(ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// fmt.Println(<-ch)

}