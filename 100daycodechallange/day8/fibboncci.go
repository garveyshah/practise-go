/*
Day 8: Write a program to generate the Fibonacci sequence up to n terms.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Input the nth term: ")

	reader := bufio.NewReader(os.Stdin)
	term, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading the input: ", err)
		return
	}

	term = strings.TrimSpace(term)

	num, err := CustomAtoi(term)
	if err != nil {
		fmt.Println("Error!! ", err)
		return
	}

	fmt.Println(num)
	fmt.Println()
	fmt.Println(Fibonacci(num))
}

// Function to generate fibonnaci numbers upto the nth term.
func Fibonacci(num int) []int {
	if num <= 0 {
		return []int{}
	}

	result := []int{1}
	a, b := 0, 1

	for i := 2; i <= num; i++ {
		c := a + b
		a = b
		b = c
		result = append(result, c)

	}
	return result
}

// Funtion to convert a string digit to an interger.
func CustomAtoi(s string) (int, error) {
	var result int

	for _, num := range s {
		if num < '0' || num > '9' {
			return 0, fmt.Errorf("invalid charachter: %c", num)
		}
		result = result*10 + int(num-'0')
	}
	return result, nil
}
