package piscine

func ToUpper(s string) string {
	result := ""
	for _, Upc := range s {
		if Upc < 'a' || Upc > 'z' {
			result += string(Upc)
		} else {
			result += string(Upc - 32)
		}
	}
	return result
}
