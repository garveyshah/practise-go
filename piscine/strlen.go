package project01

/*func StrLen(s string) int {
	str := []rune(s)
	return len(str)
}
*/

func StrLen(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}
