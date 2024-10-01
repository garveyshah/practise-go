package 

import (
	"fmt"
)

// AddDigits() returns the sum of elements in a int slice
func AddDigits(s string) (sum int) {
	var res []int
	for _, digit := range s {
		num, err := Atoi(string(digit))
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		res = append(res, num)
	}
	for _, num := range res {
		sum += num
	}
	return sum
}

// Atoi() converts a string to int
func Atoi(s string) (num int, err error) {
	var (
		sign bool
		res  int
	)

	if s == "0" {
		return 0, nil
	}

	if s[0] == '-' {
		sign = true
		s = s[1:]
	}

	for _, char := range s {
		if char < '0' || char > '9' {
			return 0, fmt.Errorf("invalid char %v", string(char))
		} else {
			res = res*10 + int(char-'0')
		}
	}
	if sign {
		res *= -1
	}
	return res, nil
}
