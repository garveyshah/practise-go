package piscine

func NRune(s string, n int) rune {
	nthRune := []rune(s)
	if n <= 0 || n > len(s) {
		return 0
	}
	return nthRune[n-1]
}
