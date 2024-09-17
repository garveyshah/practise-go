package main

import (
	"fmt"
	"math"
)

// function to calculate the slope and intercept
func linearRegression(xs, ys []float64) (slope, intercept float64) {
	// Number of data points
	n := float64(len(xs))

	// Sums required for slope and intercept formulas
	var sumX, sumY, sumXY, sumX2 float64
	for i := 0; i < len(xs); i++ {
		sumX += xs[i]
		sumY += ys[i]
		sumXY += xs[i] * ys[i]
		sumX2 += xs[i] * xs[i]
	}

	// Calculate slope (b1)
	slope = (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)

	// Calculate intercept (b0)
	intercept = (sumY - slope*sumX) / n

	return slope, intercept
}

// Function to calculate the range the next number may fall and predic the next number employing the linear regression model.
func Prectidor(x, slope, intercept float64) (float64, float64, float64) {
	guess := (slope * x) + intercept

	residual := x - guess

	return x, guess, residual
}

func CalculateRange(guessArray []float64, stdError float64) (float64, float64) {
	var upperB, lowerB float64

	for _, guess := range guessArray {
		upperB = guess - (stdError * 1)
		lowerB = guess + (stdError * 1)
	}
	return upperB, lowerB
}

// Function to calculate Variance of the residuals
func StandardError(residuals []float64) float64 {
	n := len(residuals)
	var ssd, stdError float64

	for _, num := range residuals {
		ssd += num * num
		stdError = ssd / float64(n-2)
	}
	return float64(math.Sqrt(stdError))
}

func Output(xs, ys []float64) {
	var guess, upperB, lowerB, residual, x, stdError float64
	slope, intercept := linearRegression(xs, ys)
	var residsArray, guessArray []float64

	for _, num := range xs {
		x, guess, residual = Prectidor(num, slope, intercept)

		residsArray = append(residsArray, residual)
		guessArray = append(guessArray, guess)

		stdError = StandardError(residsArray)
		upperB, lowerB = CalculateRange(guessArray, stdError)

		fmt.Printf("Guess:	%.2f\n", guess)
		fmt.Printf("\nx: %.2f\nUpperB: %.2f\nLowerB: %.2f\nResidual: %.2f\n\n", x, upperB, lowerB, residual)

	}

	// Output results
	fmt.Printf("Slope:     %.2f\n", slope)
	fmt.Printf("Intercept: %.2f\n\n", intercept)
}

func main() {
	// Sample data
	xs := []float64{1, 2, 3, 4, 5}
	ys := []float64{2, 4, 5, 4, 5}

	Output(xs, ys)
}
