package piscine

func IsNumeric(s string) bool {
	for _, B := range s {
		if B < '0' || B > '9' {
			return false
		}
	}
	return true
}
