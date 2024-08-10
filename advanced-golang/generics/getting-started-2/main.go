package main

import "fmt"

func main() {
	//1. Generic Functions
	// PrintValue(42)      // Works with int
	// PrintValue("Hello") // Works with string

	// fmt.Printf("\"Non-Generic\" - Has a %v\n", Has([]string{"a", "b"}, "a"))
	// fmt.Printf("\"Non-Generic\" - Has c %v\n", Has([]string{"a", "b"}, "c"))
	// fmt.Printf("\"Non-Generic\" - Has 2 %v\n", Has([]int{1,2,3}, 2))

	// fmt.Printf("\"Generic\" - Has a %v\n", Has([]string{"a", "b"}, "a"))
	// fmt.Printf("\"Generic\" - Has c %v\n", Has([]string{"a", "b"}, "c"))
	// fmt.Printf("\"Generic\" - Has 2 %v\n", Has([]int{1,2,3}, 2))

// // 2. Generic Constrains
// fmt.Println(Add(10, 20))          // Works with int
// fmt.Println(Add(10.5, 20.3))      // Works with float64

	// 3. Generic Methods
	// intPair := Pair[int]{first: 1, second: 2}
    // intPair.Swap()
    // fmt.Println(intPair) // Outputs: {2 1}

	// //4. Generic Structs
	// intContainer := Container[int]{value: 100}
    // stringContainer := Container[string]{value: "Go Generics"}

    // fmt.Println(intContainer.GetValue())  // Outputs: 100
    // fmt.Println(stringContainer.GetValue()) // Outputs: Go Generics
	
}

// // Non Generic Functions.

// func PrintValue(value int) {
// 	fmt.Printf("Non-Generic-func: %v\n", value)
// }

// func Has(list []string, value string) bool {
// 	for _, v := range list {
// 		if v == value {
// 			return true
// 		}
// 	}
// 	return false
// }

// 1. Generic Functions 

// func PrintValue[T any](value T) {
// 	fmt.Printf("Generic-func: %v\n", value)
// }

// T: A type parameter that can represent any type.
// any: A type constraint that indicates T can be any type.

// func Has[T comparable](list []T, value T) bool{
// 	for _, v := range list {
// 		if v == value {
// 			return true
// 		}
// 	}
// 	return false
// }


// 2. Generic Constrains
// func Add[T int | float64](a, b T) T {
//     return a + b
// }


// // 3.  Generic Methods 
// type Pair[T any] struct {
//     first, second T
// }

// func (p *Pair[T]) Swap() {
//     p.first, p.second = p.second, p.first
// }


// // 4. Generic Structs

// type Container[T any] struct {
//     value T
// }

// func (c *Container[T]) GetValue() T {
//     return c.value
// }
