package atoi

import "fmt"

func Atoi(s string) (int, error) {
	var result int
	var IsNeg bool

	if s[0] == '-' {
		IsNeg = true
		s = s[1:]
	}

	for _, char := range s {
		if char < '0' || char > '9' {
			return 0, fmt.Errorf("invalid char \"%v\"", string(char))
		}
		result = result*10 + int(char-'0')
	}
	if IsNeg {
		result *= -1
	}
	return result, nil
}
