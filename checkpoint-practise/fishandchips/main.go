package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
	"github.com/01-edu/z01"
)

func main() {
	// Print(FishAndChips(4))
	// Print(FishAndChips(9))
	// Print(FishAndChips(6))

	// fmt.Println()
	// fmt.Println(FishAndChips(4))
	// fmt.Println(FishAndChips(9))
	// fmt.Println(FishAndChips(6))

	args := []int{20, 15, -125, 10, 5, 4, -9, 9, 24, 0, -85, 6, 11}

	for i := 0; i < len(args); i++ {
		challenge.Function("FishAndChips", FishAndChips, solutions.FishAndChips, args[i])
	}
}

func FishAndChips(n int) string {
	if n < 0 {
		return "error: number is negative"
	}

	if n%3 == 0 && n%2 == 0 {
		return "fish and chips"
	}

	if n%2 == 0 {
		return "fish"
	}

	if n%3 == 0 {
		return "chips"
	} 
	
	return "error: non divisible"
}

func Print(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
	z01.PrintRune(10)
}
