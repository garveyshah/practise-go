package piscine

func Compact(ptr *[]string) int {
	Astr := *ptr
	count := 0
	for _, val := range *ptr {
		if val != "" {
			Astr[count] = val
			count++
		}
	}
	*ptr = Astr[0:count]
	return count
}
