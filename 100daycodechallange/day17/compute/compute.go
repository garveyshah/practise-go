package compute

func Fibonnaci(n int) (fib int) {
	if n < 0 {
		return 0
	}

	// if n == 1 {
	// 	return 1
	// }

	var a, b, d int = 0, 1, 0

	for i := 0; i <= n; i++ {
		c := a + b 
		b = c
		a = b
		d = c
	}
	return d
}