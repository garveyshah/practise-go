/* Day 4: Write a program to find the greatest common divisor (GCD) of two numbers.*/

// Method 1: The LCM Method
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Input the Numbers to use: ")

	reader := bufio.NewReader(os.Stdin)
	nums, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error!! Reading Input")
		return
	}

	nums1 := strings.TrimSpace(nums)
	nums2 := strings.Split(nums1, " ")
	var result []int

	for _, num := range nums2 {
		num1, err := CustomAtoi(string(num))
		if err != nil {
			fmt.Println("Invalid Input:", num)
			return
		}
		result = append(result, num1)
	}
	fmt.Println("\n array", result)

	gcd := GCD(result)

	fmt.Println("\n GCD is", gcd)
}

// func GCD calculates the gcd range of any numbers.
func GCD(S []int) int {
	if len(S) == 2 {
		return gcdpair(S[0], S[1])
	}
	result := S[0]
	for i := 1; i < len(S); i++ {
		result = gcdpair(result, S[i])
	}
	return result
}

// func gcdpair calculates the gcd of any 2 numbers.
func gcdpair(a, b int) int {
	var gcd int
	lcm := LCM(a, b)
	product := a * b
	gcd = product / lcm
	return gcd
}

// calculates the Lcm of a range of numbers.
func LCM(a int, b int) int {
	result := []int{}

	seen := make(map[int]bool)
	Ms1 := Multiples(a)
	Ms2 := Multiples(b)

	for _, nums := range Ms1 {
		seen[nums] = true
	}

	for _, val := range Ms2 {
		if seen[val] {
			result = append(result, val)
			delete(seen, val)
		}
	}
	result = BubbleSort(result)
	return result[0]
}

// func Multiples genarates sorted multiples of a number.
func Multiples(num int) []int {
	var multiples []int

	for i := num; i < 10000; i += num {
		multiples = append(multiples, i)
	}
	multiples = BubbleSort(multiples)
	return multiples
}

// func CustomAtoi converts the numbers provided as strings to integers
func CustomAtoi(s string) (int, error) {
	var result int
	for _, num := range s {
		if num < '0' || num > '9' {
			return 0, fmt.Errorf("invalid character: %c", num)
		}
		result = result*10 + int(num-'0')
	}
	return result, nil
}

// Func BubbleSort Sorts a slice of intergers.
func BubbleSort(s []int) []int {
	n := len(s)
	if n < 1 {
		return []int{}
	}

	for i := 0; i < n-1; i++ {
		for j := i; j < n-1-i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	return s
}
