package conv

import "fmt"

// Atoi() converts strings into intergers
func Atoi(s string) (num int, err error) {
	if s == "" {
		return 0, fmt.Errorf("no input")
	}

	var sign bool
	if s[0] == '-' {
		sign = true
		s = s[1:]
	}

	for _, char := range s {
		if char < '0' || char > '9' {
			return 0, fmt.Errorf("invalid character %v", string(char))
		} else {
			num = num*10 + int(char-'0')
		}
	}

	if sign {
		num *= -1
	}

	return num, nil
}
