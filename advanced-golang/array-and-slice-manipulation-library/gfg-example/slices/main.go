package main

import "fmt"

func main() {
	//A demonstration of how to create a slice in go
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:4]

	fmt.Println("Array: ", array)
	fmt.Println("Slice: ", slice)
	fmt.Println()
	
	// Demonstration of how to add elements to slice in Go
	slice1 := []int{1,2,3}
	slice1 = append(slice1, 4,5,6)

	fmt.Println("Slice: ", slice1)
}
