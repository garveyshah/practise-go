// Goroutine is an idependent function that excutes simultanuosly in some separete lightweight threads managed by go.

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go helloworld(&wg)
	go goodbye(&wg)
	wg.Wait()

	// time.Sleep( 32 * time.Microsecond)
}

// prints "Hellow world"
func helloworld(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hellow World")
}

// prints "Good Bye"
func goodbye(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Good Bye")
}

// The ouput is the same as the previous one,
// but it doesn't block the main for seconds

// 1. `wg.Add(int): This method indicates the number of goroutines to wait.
// In the above code, I have provided 2 for 2 different goroutines.
// Hence the internal counter wait becomes 2`
// 2. wg.Wait(): This method blocks the execution of code until the internal counter 
// becomes 0
// wg.DOne(): This will reduce the internal counter value bby 1.

// NOTE: IF a wait goup is explisitly assed into functions, it should be added by a pointer.

