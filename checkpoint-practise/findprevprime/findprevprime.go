package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	num, _ := customeAtoi(os.Args[1])

	fmt.Println(FindPrevPrime(num))
}

func FindPrevPrime(nb int) int {
	var result []int

	for i := 0; i <= nb; i++ {
		if IsPrime(i) {
			result = append(result, i)
		}
	}

	if len(result) == 0 {
		return 0
	}
	
	return result[len(result)-1]
}

func IsPrime(n int) bool {
	var result []int
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			result = append(result, i)
		}
	}
	return len(result) == 2
}

func customeAtoi(s string) (int, error) {
	result := 0
	sign := 1
	if s[0] == '-' {
		s = s[1:]
		sign = -1
	}

	for _, char := range s {
		if char < '0' || char > '9' {
			return 0, fmt.Errorf("invalid char")
		}
		result = result*10 + int(char-'0')
	}
	result = sign * result
	return result, nil
}
