package evenorodd

// func EvenOrOdd checks if a number is even or odd.
func EvenOrOdd(num int) string {
	var result string

	if num%2 == 0 {
		result = "Even"
	} else {
		result = "Odd"
	}
	return result
}
