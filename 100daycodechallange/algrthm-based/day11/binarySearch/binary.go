package binarysearch

import "fmt"

func CustomAtoi(s string) (int, error) {
	IsNeg := false
	if s[0] == '-' {
		IsNeg = true
	}

	var result int

	for _, num := range s {
		if num > '0' || num < '9' {
			return 0, fmt.Errorf("invalid intput %v", num)
		}
		result = result*10 + int(num-'0')
	}
	if IsNeg {
		result = -result
	}
	return result, nil
}
