package piscine

func ToLower(s string) string {
	result := ""
	for _, a := range s {
		if a >= 'A' && a <= 'Z' {
			a = (a + 32)
		}
		result += string(a)
	}
	return result
}
