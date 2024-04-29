package piscine

func Unmatch(a []int) int {
	for _, num := range a {
		i := 0
		for _, value := range a {
			if value == num {
				i++
			}
		}
		if i == 1 || i%2 == 1 {
			return num
		}
	}
	return -1
}
