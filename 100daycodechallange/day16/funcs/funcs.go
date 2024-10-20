package funcs

// Merges 2 arrays and calls Sort function to sort its return elements in ascending order
func Merge(a, b []int) (c []int) {
	n := 0
	for i := 0; i < len(a); i++ {
		c = append(c, a[i])
		c = append(c, b[n])
		n++
	}
	 Sort(c)
	return c
}

// Sorts an array in ascending order
func Sort(input []int) []int {
	n := len(input)
	if n < 1 {
		return []int{}
	}

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
			}
		}
	}
	return input
}
