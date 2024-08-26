package main

import "fmt"

func main() {
	// make new
	// new - only allocates - no initializtion of memory
	// make - allocate and initiolization of meomory - non zeroed storage

	// var score map[string]int
	score := make(map[string]int)
	score["Godwin"] = 10
	fmt.Println(score)

	score["Godwin"] = 10
	score["Adda"] = 20
	score["odwin"] = 30
	score["win"] = 180
	score["God"] = 100
	fmt.Println(score, "\n")

	getWinScore := score["win"]
	fmt.Println("win:", getWinScore)

	// delete
	delete(score, "win")
	fmt.Println(score, "\n")

	// looping

	// k is variable to store all the keys
	// v is a variable to store all the values
	// when only one

	for k, v := range score {
		fmt.Println("key: ", k, "value: ", v)
	}
}
