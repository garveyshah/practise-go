package project01

func BubbleSort(numbers []int, a, b int) {
	//swap, check location, bubble sort round one, repeat sort.
	temp := numbers[a]
	numbers[a] = numbers[b]
	numbers[b] = temp
	for j := 0; j < len(numbers); j++ {
		for i := 0; i < len(numbers)-1; i++ {
			if numbers[i] > numbers[i+1] {
				BubbleSort(numbers, i, i+1)
			}
		}

	}
}
