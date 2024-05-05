package sorting

func BubbleSort(input []int) []int {
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
