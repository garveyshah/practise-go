package main

import "fmt"

func main() {
	fmt.Println(Rev([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(Rev([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(Rev([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	fmt.Println(Rev([]int{1, 2, 3}, []int{}))
}

func Rev(slice1, slice2 []int) (result []int) {
	len1, len2 := len(slice1), len(slice2)

	i, j := len1-1, len2-1

	for i >= 0 && j >= 0 {
		if len1 > len2 {
			result = append(result, slice1[i])
			len1--
			i--
		} else if len2 >= len1 {
			result = append(result, slice2[j])
			len2--
			j--
		} else {
			result = append(result, slice1[i])
			len2--
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
