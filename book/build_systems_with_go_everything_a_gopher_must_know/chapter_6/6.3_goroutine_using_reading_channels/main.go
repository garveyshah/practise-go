/*
	In example below `make(chan int)` instantiates a channel that can send or receive integers.
	In this particular example, the `generator` functions runs in a goroutine that computes the sum of the first five intergers and sends it through the `channel`.
	Meanwhile, the main function waits until something is sent through the channel with `result := <- ch`.
	Notice that this last operation is blocking and will nit be completed until something is sent through the channel. 
*/
package main

import (
	"fmt"
	"time"
)

func generator(ch chan int) {
	sum := 0

	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 500)

		sum =  sum + i
	}

	ch <- sum
}

func main() {
	ch := make(chan int)

	go generator(ch)

	fmt.Println("main waits for result...")
	result := <- ch

	fmt.Println(result)
}

// We can enhance this example using the channel in both directions: reading and writing