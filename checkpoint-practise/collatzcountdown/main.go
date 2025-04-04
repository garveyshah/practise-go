package main

import "fmt"

func CollatzCountdown(start int) (steps int) {
	if start <= 0 {
		return -1
	}

	for start != 0 {
		if start %2 == 0 {
			start = start / 2
		} else {
			start = 3 * start + 1
		}
		steps ++
	}
	return steps
}

func main() {
	steps := CollatzCountdown(12)
	fmt.Println(steps)
}