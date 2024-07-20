/*
 Prime Number Generator - Generate all prime numbers up to n.
*/

package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		PrintString("Usage: go run . [number]")
		return
	}

	num, err := CustomAtoi(os.Args[1])
	if err != nil {
		PrintString("Error " + err.Error())
		return
	}

	result := PrimeNumberGenerator(num)
	fmt.Println(result)
}

// Function PrimeNumberGenerator generates prime numbers up to a given number.
func PrimeNumberGenerator(n int) []int {
	var primes []int

	for i := 2; i <= n; i++ {
		if IsPrime(i) {
		primes = append(primes ,i)
		}
	}
	return primes
}

// Function IsPrime checks if an integer is a prime number or not.
func IsPrime(n int) bool {
	var factors []int

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			factors = append(factors, i)
		}
	}

	return len(factors) == 2
}

// Function CustomAtoi converts string digit to an interger.
func CustomAtoi(s string) (int, error) {
	var result int
	var sign bool

	if len(s) == 0 {
		return 0, errors.New("empty string")
	}

	if s[0] == '-' {
		sign = true
		s = s[1:]
	}

	for _, char := range s {
		if char < '0' || char > '9' {
			return 0, errors.New("invalid character")
		}
		result = result*10 + int(char-'0')
	}
	if sign {
		result *= -1
	}
	return result, nil
}

// Function PrintString a String rune by rune.
func PrintString(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
	z01.PrintRune(10)
}
