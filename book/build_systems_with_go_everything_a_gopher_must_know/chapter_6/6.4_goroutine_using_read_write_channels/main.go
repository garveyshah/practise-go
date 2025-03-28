/*
	In the example `generator` receives the number of elements to iterate through the channel.
	This function will be blocked until the number is received.
	You can observe this by manipulating the sleep time in the main function before sending 'n' value through
	the channel
*/

package main

import (
	"fmt"
	"time"
)

func generator(ch chan int) {
	fmt.Println("generator waits for n ")

	n := <- ch

	fmt.Println("n is", n)

	sum := 0 
	for i := 0; i < n; i++ {
		sum = sum + i
	}

	ch <- sum
}

func main() {
	ch := make(chan int)

	go generator(ch)

	fmt.Println("main waits for result ...")

	time.Sleep(time.Second * 10)

	ch <- 5

	result := <- ch

	fmt.Println(result)
}

/*
	Channels can be used to read or write. However, observe that the arrow statement <- always goes from the right to the left.

	`
	ch <- 5 // send 5 through channel
	n := <- ch // initialize n with value from  channel
	<- ch // wait until something is sent through ch
	`
*/

