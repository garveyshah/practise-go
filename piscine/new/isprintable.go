package piscine

func IsPrintable(s string) bool {
	for _, P := range s {
		if P < 32 || P > 126 {
			return false
		}
	}
	return true
}
