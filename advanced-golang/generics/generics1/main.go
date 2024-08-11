package main

import "fmt"

func main() {
	result := Add(1, 2)
	fmt.Printf("result: %+v\n",result)
}

func Add[T int | float64](a T, b T) T{
	return a + b
}

// Type Parameters.
