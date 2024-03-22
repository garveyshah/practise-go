package piscine

func LastRune(s string) rune {
	lRune := []rune(s)
	return lRune[len(s)-1]
}
