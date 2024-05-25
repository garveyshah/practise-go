package main

import (
	"fmt"
)

func main () {
	b := []int{1,1,3,2,2}

	fmt.Println(Unmatch(b))
}

func Unmatch(v []int) int{
	count := 0
	var char2 int
	for _, char := range v {
		for _, next := range v {
			if char == next {
				count++
			} 
		}
		if count%2 == 0 {
			char2 = -1
		} else {
		char2 = char
	}
	}	
	return char2
}

