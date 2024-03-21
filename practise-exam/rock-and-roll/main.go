package main

import "fmt"

func RockAndRoll(nb int) string {
	if nb < 0 {
		return "error: number is negative\n"
	}
	isDivisible3 := nb%3 == 0
	isDivisible2 := nb%2 == 0

	if isDivisible3 && isDivisible2 {
		return "rock and roll\n"
	} else if isDivisible2 {
		return "rock\n"
	} else if isDivisible3 {
		return "roll\n"
	}
	return "error: non divisible\n"
}

func main() {
	fmt.Println(RockAndRoll(4))
	fmt.Println(RockAndRoll(9))
	fmt.Println(RockAndRoll(6))
}
