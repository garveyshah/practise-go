/*
Day 9: Implement a function to find the factorial of a number.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Enter the number... Then press enter.")

	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading in put", err)
		return
	}
	str = strings.TrimSpace(str)

	num, err := CustomAtoi(str)
	if err != nil {
		fmt.Println("Error!!", err)
		return
	}
	fmt.Println(num)

	fact := Factorial(num)
	fmt.Printf("\nthe factorial of %d is %d\n", num, fact)
}

func Factorial(num int) int {
	if num < 0 || num > 20 {
		return 0
	}
	if num == 0 || num == 1 {
		return 1
	}
	result := 1
	for i := 1; i <= num; i++ {
		result = result * i
	}
	return result
}

// Fuction to convert string digits to intergers.
func CustomAtoi(s string) (int, error) {
	result := 0
	for _, char := range s {
		if char < '0' || char > '9' {
			return 0, fmt.Errorf("invalid input:\"%v\"", string(char))
		}
		result = result*10 + int(char-'0')
	}

	return result, nil
}
