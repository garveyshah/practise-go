package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := append(random.IntSliceBetween(0, 99999), 5, 4, 1, 0)
	for _, arg := range args {
		challenge.Function("FindPrevPrime", FindPrevPrime, solutions.FindPrevPrime, arg)
	}
}

func FindPrevPrime(n int) int {
	for n > 1 {
		if IsPrime(n) {
			return n 
		}
		n-- 
	}
	if n <= 1 {
		return 0 
	}
	return n
}

func IsPrime(num int) bool {
	if num < 1 {
		return false 
	}

	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

