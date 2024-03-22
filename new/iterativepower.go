package piscine

func IterativePower(nb int, power int) int {
	for power < 0 {
		return 0
	}
	result := 1
	for i := 1; i <= power; i++ {
		result = nb * result
	}
	return result
}
