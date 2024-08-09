package main

import "fmt"

func main() {
	fmt.Printf("Has a %v\n", Has([]string{"a", "b"}, "a"))
	fmt.Printf("Has c %v\n", Has([]string{"a", "b"}, "c"))
	fmt.Printf("Has 2 %v\n", Has([]int{1,2,3}, 2))
}

// // Non Generic
// func Has(list []string, value string) bool {
// 	for _, v := range list {
// 		if v == value {
// 			return true
// 		}
// 	}
// 	return false
// }

// Generic
func Has[T comparable](list []T, value T) bool{
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}
