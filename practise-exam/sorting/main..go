package main

import (
	"fmt"
)

func BubbleSort(numbers []int, a, b int) {
	// ( 2 1 ) -> ( 1 2 )
	//( 2 1 )	-> ( 1 1)
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

func main() {
	// swap, check order, one-round bubbling, repeated n times
	numbers := []int{98, 5, 865, 1000, 100, 907, 67, 9876, 78, 45, 87, 7382, 282}
	BubbleSort(numbers, 0, 1)
	fmt.Println("bubble sorted =>", numbers)
}
