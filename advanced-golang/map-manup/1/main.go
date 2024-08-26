package main

import "fmt"

func main() {
	scores := map[string]int{"Alice": 90, "Bob": 85, "Char": 88}
	fmt.Println(Read("Alice", scores), "\n")

	PrintValue(scores)
	fmt.Println()
	KeyOrValue(scores)
}

func Read(key string, scores map[string]int) int {
	return scores[key]
}

// Using the key word
func PrintValue(score map[string]int) {
	for key, value := range score {
		fmt.Println("Name: ", key, "Score: ", value)
	}
}

// Iterating Over Keys or values Only:
func KeyOrValue(scores map[string]int) {
	// only the keys of thw map are iterated over
	for key := range scores {
		fmt.Println("Name: ", key)
	}

	// Only the values are iterated over. the '_' symbol is used 
	// to discard the key varaible iterating over the values
	for _, value := range scores {
		fmt.Println("Scores: ", value)
	}
}
