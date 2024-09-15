package main

import (
	"fmt"
)

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

func main() {
	// Sample data
	xs := []float64{1, 2, 3, 4, 5}
	ys := []float64{2, 4, 5, 4, 5}

	// Perform linear regression
	slope, intercept := linearRegression(xs, ys)

	// Output results
	fmt.Printf("Slope: %.2f\n", slope)
	fmt.Printf("Intercept: %.2f\n", intercept)
}
