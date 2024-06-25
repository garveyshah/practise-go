/*
Day 4: Write a program to find the greatest common divisor (GCD) of two numbers.
*/

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
		fmt.Println("Error!! Readind Input")
		return
	}

	nums1 := strings.TrimSpace(nums)
	nums2 := strings.Split(nums1, " ")
	var result []int

	fmt.Println("Nums2", nums2)

	for _, num := range nums2 {
		num1, _ := CustomAtoi(string(num))
		result = append(result, num1)
	}
	fmt.Println("\n array", result)

	gcd := GCD(result)

	fmt.Println("\n GCD is", gcd)
}

func GCD(S []int) int {
	var gcd int

	if len(S) == 2 {
		gcd = gcdpair(S[0], S[1])
	} else {
		result := S[0]
		for i := 1 ; i < len(S); i++ {
			result = gcdpair(result, S[i])
		}
		gcd = result
	}
	return gcd
}

func gcdpair(a, b int) int {
	var gcd int
	lcm := LCM(a, b)
	product := a * b
	gcd = product / lcm
	return gcd
}


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

func Multiples(num int) []int {
	var multiples []int

	for i := num; i < 1000; i += num {
		multiples = append(multiples, i)
	}
	multiples = BubbleSort(multiples)
	return multiples
}

func CustomAtoi(s string) (int, error) {
	var result int
	for _, num := range s {
		if num < '0' || num > '9' {
			return 0, fmt.Errorf("invalid chaaracter: %c", num)
		}
		result = result*10 + int(num-'0')
	}
	return result, nil
}

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
