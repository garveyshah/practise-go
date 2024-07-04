/*
Day 10: Write a program to check if a number is a prime number.
*/

// Is PrimeFUCTIONS

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter the number....")

	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error! reading input", err)
	}

	str = strings.TrimSpace(str)

	num, err := CustomAtoi(str)
	if err != nil {
		fmt.Println("Error!! ..", err)
		return
	}
	if num <= 0 {
		fmt.Println("Error!! .. Enter a positive integer greater than 0")
		return
	}
	str2, slice := IsPrime(num)
	if len(slice) >= 2 {
		fmt.Printf("\n\"%d\" %v.\nIt has %d factors :-> %v\n.", num, str2, len(slice), slice)
	} else {
		fmt.Printf("\n\"%d\" %v.\nIt has %d factor :-> %v\n.", num, str2, len(slice), slice)
	}
}

func CustomAtoi(s string) (int, error) {
	if len(s) == 0 {
		return 0, fmt.Errorf("invalid input: empty string")
	}
	var result int
	isNegative := false
	var start int

	if s[0] == '-' {
		isNegative = true
		start = 1
	}

	for i := start; i < len(s); i++ {
		num := s[i]
		if num < '0' || num > '9' {
			return 0, fmt.Errorf("invalid input \"%v\"", num)
		}
		result = result*10 + int(num-'0')

	}
	if isNegative {
		result = -result
	}

	return result, nil
}

func IsPrime(num int) (string, []int) {
	if num < 2 {
		return "Is NOT A PRIME NUMBER,It is less than 0", []int{num}
	}

	resultSlice := []int{}
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			resultSlice = append(resultSlice, i)
		}
	}

	isPrime := len(resultSlice) == 2

	if isPrime {
		return "Is a PRIME NUMBER", resultSlice
	}
	return "Is NOT A PRIME NUMBER", resultSlice
}
