package main

import "fmt"

// swap, check order, one-round bubbling, repeated n times
func Swap(number []int, b, c int) {

	temp := number[b]
	number[b] = number[c]
	number[c] = temp

}

func main() {
	number := []int{6, 1, 4, 2, 9}
	Swap(number, 0, 1)
	fmt.Println(number)
}
