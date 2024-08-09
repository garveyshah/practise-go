package main

import (
	"fmt"
)

func main() {
	fmt.Println(Gcd(42, 10))
	fmt.Println(Gcd(42, 12))
	fmt.Println(Gcd(14, 77))
	fmt.Println(Gcd(17, 3))
}

func Gcd(a, b uint) uint {
	if a == 0 || b == 0 {
		return 0
	}

	af := Factors(a)
	bf := Factors(b)
	NumMap := make(map[uint]bool)

	for _, char := range af {
		NumMap[char] = true
	}
	var new []uint
	for _, seen := range bf {
		if _, ok := NumMap[seen]; ok {
			new = append(new, seen)
		}
	}
	if len(new) > 0 {
		return new[len(new)-1]
	}
	return 0
}

func Factors(a uint) []uint {
	var result []uint
	b := int(a)
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			result = append(result, uint(i))
		}
	}
	return result
}
