# CHAPTER 6 CONCURRENCY

Go was designed with the simplicity in mind. It facilitates the creation of concurrent programs, inter-thread communication, and other topics that in other languages require deep knowledge. 
At the end of this Chapter we will understand concepts such as goroutines, channels, and how they can be used to design sophisticated concurrent programs.

# 6.1 GOROUTINES

A Goroutine is a lightweight thread managed by the Go runtime. 
Goroutines are declared using the `go` statement followed by the function to be executed.

Goroutines are very lightweight with a very small memory demand (only 2Kb) when compared with a thread.  We can expect to have several goroutines concurrently running. This can as easy as invoking the `go` statement when required.


# 6.2 CHANNELS
Channels are a mechanism that provides communication for concurrently running functions.
A channel can send or receive elements of a specified type.

Channels are instantiated using the `make` built-in function

# 6.2.1 Buffered channels
Channels can be buffered or unbuffered. The statement `make(chan int)` generates an unbuffered channel. 
Unbuffered channels have no capacity, this means that both sides of the channel must be ready to send or receive data.
On the other side, buffered channels can be declared with `make(chan int, 10)` where `10` is the buffer size. 
In this case, the channel can store values independently of the readines of sender and receiver.