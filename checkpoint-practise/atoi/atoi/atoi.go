package atoi

func Atoi(s string) int {
	var result int
	var IsNeg bool

	if s[0] == '-' {
		IsNeg = true
		s = s[1:]
	}

	for _, char := range s {
		if char < '0' || char > '9' {
			return 0
		}
		result = result*10 + int(char-'0')
	}
	if IsNeg {
		result *= -1
	}
	return result
}
