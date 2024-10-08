package calc

func Power(num, exp int) (power int) {
	if exp == 0 {
		return 1
	}
	for i := 1; i < exp; i++ {
		power = power * num
	}
	return power * power
}
