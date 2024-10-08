package calc

func Power(num, exp int) (power int) {
	if exp == 0 {
		return 1
	}
	power = 1
	for i := 1; i <= exp; i++ {
		power = power * num
	}
	return power
}

func Factors(n int) []int {
	factors := []int{}
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}
// funct primeFa