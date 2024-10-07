package checker

func Prefix(s []string) string {
	shortest := sort(s)

	for i := 0; i < len(shortest); i++ {
		for _, word := range s {
			if shortest[i] != word[i] {
				return shortest[:i]
			}
		}
	}

	return shortest
}

func sort(s []string) string {
	shortest := s[0]
	for _, word := range s {
		if len(word) < len(shortest) {
			shortest = word
		}
	}
	return shortest
}
