/*
    The example below creates three goroutines that print a number after sleeping for a given time.
    The output shows the proportion of the messages we expect depending on the sleeping time.
    Observe that multiple executions of this program may not return the same output. 
    Why? Because the goroutines initialization time may vary and because we using the console ouput which may is a single ouput for multiple routines.
*/

package main

import (
    "fmt"
    "time"
)

func ShowIt(t time.Duration, num int) {
    for {
        time.Sleep(t)
        fmt.Println(num)
    }
}

func main() {
    go ShowIt(time.Second, 100)
    go ShowIt(time.Millisecond*500, 10)
    go ShowIt(time.Millisecond*250, 1)
    
    time.Sleep(time.Millisecond*1200)
}
