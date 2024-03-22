package piscine

func Join(strs []string, sep string) string {
	var result string
	for i, char := range strs {
		result += char
		if i != len(strs)-1 {
			result += sep
		}
	}
	return result
}
