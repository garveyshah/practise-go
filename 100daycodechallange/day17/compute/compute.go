package compute

func Fibonnaci(n int) int {
	if n <= 1 {
		return n
	}

	return Fibonnaci(n-1) + Fibonnaci(n-2)
}