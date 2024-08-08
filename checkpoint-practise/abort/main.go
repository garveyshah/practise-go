package main

import (
	"fmt"
)

func main() {
	middle := Abort(2, 3, 8, 5, 7)
	fmt.Println(middle)
}

func Abort(a, b, c, d, e int) int {
	slice := BubbleSort([]int{a, b, c, d, e})
	var result int
	for i := 0; i <= len(slice)-1; i++ {
		if len(slice)%2 != 0 {
			result = slice[len(slice)/2]
		} else {
			result = slice[(len(slice)/2)+1] + slice[(len(slice)/2)]
		}
	}
	return result
}

func BubbleSort(slice []int) []int {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	return slice
}
