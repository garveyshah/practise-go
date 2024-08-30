package main

import "fmt"

type Person struct {
	Name  string
	Likes []string
}

func main() {
	// Create some people with their likes
	people := []*Person{
		{Name: "Alice", Likes: []string{"cheese", "bacon"}},
		{Name: "Bob", Likes: []string{"lettuce", "bacon"}},
		{Name: "Charlie", Likes: []string{"cheese"}},
	}

	// var people []*Person
	// Create a map to group people by their likes
	likes := make(map[string][]*Person)
	for _, p := range people {
		for _, l := range p.Likes {
			likes[l] = append(likes[l], p)
		}
	}
	// Print a list of people who like cheese
	for _, p := range likes["cheese"] {
		fmt.Println(p.Name, "likes cheese")
	}

	// Print the number of people who like bacon
	fmt.Println(len(likes["bacon"]), "people like bacon")
}
