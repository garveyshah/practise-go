package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	a := 0
	for b := 0; b < 4; b++ {
		fmt.Printf("Player %v: %v, %v, %v\n", b+1, deck[a], deck[a+1], deck[a+2])
		a = a + 1
	}
}
