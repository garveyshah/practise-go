package main

import "fmt"

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

func main() {
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{}))
}

func RevConcatAlternate(slice1, slice2 []int) []int {
	var result []int
	if len(slice1) == 0 && len(slice2) != 0 {
		result = Reverser(slice2)
		return result
	} else if len(slice1) != 0 && len(slice2) == 0 {
		result = Reverser(slice1)
		return result
	} else if len(slice1) != 0 && len(slice2) != 0 {
		slice2 = Reverser(slice2)
		slice1 = Reverser(slice1)


		switch {
		case len(slice1) > len(slice2) :
				for i := len(slice2) - 1; i >= 0; i-- {
			}
		case len(slice2) > len(slice2):
			
		}
			// for j := len(slice1) - 1; j >= 0; j-- {
				result = append(result, slice1[i])
				result = append(result, slice2[i])
			// }
		return result
	} else if len(slice1) == 0 && len(slice2) == 0 {
		result = []int{}
	}
	return result
}


func Reverser(slice []int) (result []int) {
		for i := len(slice) - 1; i >= 0; i-- {
			result = append(result, slice[i])
		}
		return result
}