// Define a Go structure Car with fields Brand, Model, and Year.
// Write a program that reads car data from a CSV file,
// finds the newest car, and prints its details.
package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// Car structure
type Car struct {
	Brand string
	Model string
	Year  int
}

func main() {
	// Open the CSV file
	file, err := os.Open("Cars.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the CSV data
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	// Initialize variables to keep track of the newest car
	var newestCar Car
	newestYear := 0

	// Skip the header row
	records = records[1:]

	// Iterate through the records
	for _, record := range records {
		// Parse the year from the record
		year, err := strconv.Atoi(record[2])
		if err != nil {
			fmt.Println("Error parsing year:", err)
			continue
		}

		// Create a new Car instance
		car := Car{
			Brand: record[0],
			Model: record[1],
			Year:  year,
		}

		// Check if the current car is newer than the newest car found so far
		if car.Year > newestYear {
			newestCar = car
			newestYear = car.Year
		}
	}

	// Print the details of the newest car
	fmt.Println("Newest Car:")
	fmt.Println("Brand:", newestCar.Brand)
	fmt.Println("Model:", newestCar.Model)
	fmt.Println("Year:", newestCar.Year)
}
