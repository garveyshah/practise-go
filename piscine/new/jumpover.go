package piscine

func JumpOver(str string) string {
	Astr := ""
	if len(str) > 3 {
		for a, char := range str {
			if a%3 == 2 {
				Astr = Astr + string(char)
			}
		}
	}
	return Astr + "\n"
}
