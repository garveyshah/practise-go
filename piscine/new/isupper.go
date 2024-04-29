package piscine

func IsUpper(s string) bool {
	for _, char := range s {
		if char < 'A' || char > 'z' {
			return false
		}
		if char >= 'a' || char > 'z' {
			return false
		}
	}
	return true
}
