package main

import "fmt"

var planetOrbitalPeriods = map[string]float64{
	"Earth":   1.0,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

func main() {
	// Example: Calculate dog years for a given planet and age
	age := 10.0 // Your age in Earth years
	planet := "Mars"
	dogYears := dogYears(planet, age)
	fmt.Printf("Your age in dog years on %s is: %.2f\n", planet, dogYears)
}

func dogYears(planet string, age float64) float64 {
	if orbitalPeriod, ok := planetOrbitalPeriods[planet]; ok {
		return age * 7 / orbitalPeriod
	} else {
		fmt.Println("Unknown planet")
		return 0
	}
}
