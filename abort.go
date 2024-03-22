package piscine

func Abort(a, b, c, d, e int) int {
	k := []int{a, b, c, d, e}
	for i := 0; i <= len(k)-1; i++ {
		for j := i + 1; j <= len(k)-1; j++ {
			if k[i] > k[j] {
				k[i], k[j] = k[j], k[i]
			}
		}
	}
	return k[2]
}
