package piscine

func StrRev(s string) string {
	var strrev string
	for i := len(s) - 1; i >= 0; i-- {
		strrev += string(s[i])
	}
	return strrev
}
