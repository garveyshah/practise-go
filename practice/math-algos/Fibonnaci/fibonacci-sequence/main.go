/*
 Fibonacci Sequence - Generate the first n numbers of the Fibonacci sequence using a recursive function.
*/

package main

import (
	"fmt"
	"os"
)

// Recursive function to compute the nth Fibonacci number
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// Function to generate the first n Fibonacci numbers
func generateFibonacci(n int) []int {
	sequence := make([]int, n)
	for i := 0; i < n; i++ {
		sequence[i] = fibonacci(i)
	}
	return sequence
}

func CustomAtoi(str string) (int, error) {
	var result int
	var sign bool

	if str[0] == '-' {
		sign = true
		str = str[1:]
	}

	for _, char := range str {
		if char < '0' || char > '9' {
			return 0, fmt.Errorf("invalid character \"%v\"", string(char))
		}
		result = result*10 + int(char-'0')
	}
	if sign {
		result *= -1
	}
	return result, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Print("Enter the number of Fibonacci numbers to generate: ")
		return
	}
	n, err := CustomAtoi(os.Args[1])
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	if n <= 0 {
		fmt.Println("Please enter a positive integer.")
		return
	}

	fibonacciSequence := generateFibonacci(n)
	fmt.Println("Fibonacci sequence:", fibonacciSequence)
}
