package main

// var testCases = []struct {
// 	args [][]int
// 	want []int
// }{
// 	{
// 		[][]int{{1, 2, 3}, {4, 5, 6}},
// 		[]int{3, 6, 2, 5, 1, 4},
// 	},
// 	{
// 		[][]int{{4, 5, 6, 7, 8, 9}, {1, 2, 3}},
// 		[]int{9, 8, 7, 6, 3, 5, 2, 4, 1},
// 	},
// 	{
// 		[][]int{{1, 2, 3}, {4, 5, 6, 7, 8, 9}},
// 		[]int{9, 8, 7, 3, 6, 2, 5, 1, 4},
// 	},
// 	{
// 		[][]int{{1, 2, 3}, {4, 5}},
// 		[]int{3, 2, 5, 1, 4},
// 	},
// 	{
// 		[][]int{{}, {4, 5, 6}},
// 		[]int{6, 5, 4},
// 	},
// 	{
// 		[][]int{{1, 2, 3}, {}},
// 		[]int{3, 2, 1},
// 	},
// 	{
// 		[][]int{{}, {}},

// 		[]int{},
// 	},
// 	{
// 		[][]int{{1, 2, 4}, {10, 20, 30, 40, 50}},

// 		[]int{50, 40, 4, 30, 2, 20, 1, 10},
// 	},
// }

// func main() {
// 	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
// 	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
// 	fmt.Println(RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
// 	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{}))
// }

func RevConcatAlternate(slice1, slice2 []int) []int {
	var result []int

	len1, len2 := len(slice1), len(slice2)

	i, j := len1-1, len2-1

	for i >= 0 && j >= 0 {
		if len1 > len2 {
			result = append(result, slice1[i])
			len1--
			i--
		} else if len1 < len2 {
			result = append(result, slice2[j])
			len2--
			j--
		} else {
			result = append(result, slice1[i])
			len1--
			i--
		}
	}

	for i >= 0 {
		result = append(result, slice1[i])
		i--
	}
	for j >= 0 {
		result = append(result, slice2[j])
		j--
	}
	return result
}
