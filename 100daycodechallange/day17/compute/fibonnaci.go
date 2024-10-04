package compute

// Fibonnaci() calculates the nth fibonnaci number using recursion
func Fibonnaci(n int) int {
	if n <= 1 {
		return n
	}

	return Fibonnaci(n-1) + Fibonnaci(n-2)
}