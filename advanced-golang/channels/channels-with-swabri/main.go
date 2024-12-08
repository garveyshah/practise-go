package main

import "sync"

type Task struct {
	ID int
}

func main() {
	wg := new(sync.WaitGroup) // new() always returns a pointer to the data type 
	NumTask := 10
	workers := 3

	taskChan := make(chan Task, NumTask)
	resultChan := make(chan string, NumTask)
}
