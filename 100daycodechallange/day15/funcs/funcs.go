package funcs

import "fmt"

// Atoi() converts a string of digits to interger
func Atoi(yearS string) (yearN int, err error) {
	if yearS == "0" {
		return 0, nil
	}

	var sign bool
	if yearS[0] == '-' {
		sign = true
		yearS = yearS[1:]
	}

	for _, char := range yearS {
		if char < '0' || char > '9' {
			return 0, fmt.Errorf("invalid character %q", string(char))
		} else {
			yearN = yearN*10 + int(char-'0')
		}
	}

	if sign {
		yearN *= -1
	}
	return yearN, nil
}

// IsLeapYear() checks if a year is a leap year or not
func IsLeapYear(year int) (is bool) {
	return is
}
