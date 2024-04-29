package main

import "fmt"

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(IsPrime(97))
	fmt.Println(IsPrime(169))
}
