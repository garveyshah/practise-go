package main

import "fmt"

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}

func Chunk(slice []int, num int) {
	var chunk [][]int
	if num == 0 {
		fmt.Println()
		return
	}
	if len(slice) == 0 {
		fmt.Println(chunk)
		return
	}

	for i := 0; i < len(slice); i += num {
		end := i + num
		if end > len(slice) {
			end = len(slice)
		}
		chunk = append(chunk, slice[i:end])
	}
	fmt.Println(chunk)
}
