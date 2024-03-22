package piscine

func Rot14(s string) string {
	var result string

	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			char = 'a' + (char-'a'+14)%26
		}
		if char >= 'A' && char <= 'Z' {
			char = 'A' + (char-'A'+14)%26
		}
		result += string(char)
	}
	return result
}
