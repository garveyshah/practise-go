package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := []int{20, 15, -125, 10, 5, 4, -9, 9, 24, 0, -85, 6, 11}

	for i := 0; i < len(args); i++{
		challenge.Function("FishAndChips", FishAndChips, solutions.FishAndChips, args[i])
	}
}

func FishAndChips(num int) string {
	res := ""

	if num < 0 {
		res = "error: number is negative"
	}	else if num%3 == 0 && num%2 == 0 {
		res = "fish and chips"
	} else if num%3 == 0 {
		res = "chips"
	} else if num%2 == 0 {
		res = "fish"
	} else {
		res = "error: non divisible"
	}

	return res
}