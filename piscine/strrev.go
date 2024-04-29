package project01

func StrRev(s string) string {
	characters := []rune(s)
	start := 0
	end := len(characters) - 1

	for start < end {
		characters[start], characters[end] = characters[end], characters[start]
		start++
		end--
	}
	return string(characters)
}
