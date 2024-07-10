/*
Basic Probability Calculation: Write a program that calculates the probability of drawing a specific card from a standard deck of 52 cards.
*/

package main

import "fmt"

func main() {
	
	result := ProbabilityOfDrawingCard()

	fmt.Println(RoundTo3DecimPlaces(result))
}
func ProbabilityOfDrawingCard() float64 {
	totalCards := 52
	specificCard := 1 

	probability := float64(specificCard)/float64(totalCards)
	return probability
}
func RoundTo3DecimPlaces(val float64) float64{
	factor := 10000.0
	if val < 0 {
		return float64(int(val*factor-0.5))/factor
	}
	return float64(int(val*factor+0.5))/factor
}