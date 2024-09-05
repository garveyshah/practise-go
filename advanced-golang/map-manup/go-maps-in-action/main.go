package main

import "fmt"

func main() {
	// // decalaring a map using the inbuild make function
	// m := make(map[string]int)

	// // this statement sets the key "route" to the value 66:
	// m["route"] = 66

	// // this statement retrives the value stored under the key "route" and assighns it to a new variable i:
	// i := m["route"]
	// fmt.Println("the value of \"route\"", i)

	// // If the requested key doesn't exist we get the value type's zeero value.In this case the value type is int,
	// // so the the zero value is 0:
	// j := m["root"]
	// fmt.Println("the zero value of \"root\"", j)

	// // The built len function returns on the number of items in a map:
	// n := len(m)
	// fmt.Println("len of \"m\"", n)

	// // The built in delete() function removes an entry for the map:
	// delete(m, "route")
	// fmt.Println("new length of \"m\"", len(m))
	// // Note: The delete function doesn't return anything, and will do nothing if the specifed key doesn't exist.

	// // A two-value assignment tests for the existence of a key:
	// i, ok := m["route"]
	// fmt.Println("bool to check for the presense of key \"route\" in \"m\": ", ok, ", i is the value type zero value:", i)

	// // To initialize a map with some data, use a map literal:
	// commits := map[string]int{
	// 	"rsc": 3711,
	// 	"r":   2138,
	// 	"gti": 1908,
	// 	"adg": 912,
	// }
	// fmt.Println("len \"commits\"", len(commits), ": commits' contents:", commits)

	// // The same syntax may be used to initailize an empty map, which is functionality identical to using the make function"
	// m = map[string]int{}
	// fmt.Println("m", m)

	// Exploiting zero values
	/*This example traverses a linked list of Nodes and prints their values.
	It uses a map of Node pointers to detect cycles in the list.*/

	// type Node struct {
	// 	Next  *Node
	// 	Value interface{}
	// }

	// var first *Node

	// visited := make(map[*Node]bool)
	// for n := first; n != nil; n = n.Next {
	// 	if visited[n] {
	// 		fmt.Println("Cycle detected")
	// 		break
	// 	}
	// 	visited[n] = true
	// 	fmt.Println(n.Value)
	// }

	// /*In the following example, the slice people is populated with Person values. Each Person has a Name and a slice of Likes.
	// The example creates a map to associate each like with a slice of people that like it.*/

	// type Person struct {
	// 	Name  string
	// 	Likes []string
	// }

	// var people []*Person

	// likes := make(map[string][]*Person)
	// for _, p := range people {
	// 	for _, l := range p.Likes {
	// 		likes[l] = append(likes[l], p)
	// 	}
	// }

	// // To print a list of people who  who like cheese:
	// for _, p := range likes["cheese"] {
	// 	fmt.Println(p.Name, "likes cheese")
	// }
	// fmt.Println(len(likes["beacon"]), "people like bacon.")

	// // Key Types
	// // Maps keys maybe of any type that is comparable.
	// // Comparable types are boolean, numeric, string, pointer, channel,
	// // and interface types, and structs or arrays that contain only those types.
	// // slices, maps and functions cannot be compared using == and my not be used as map keys.

	// // Struct can be used to key data by multiple dimensions. For example,
	// // this map of maps can be used to tally web page hits by country:
	// hits := make(map[string]map[string]int)

	// // Each key of thee outer map is the path to a web page with its own inner map. Each nner map key is a two-letter country code. This expression retrives the numbber of times an Australian has loaded the documentation page:
	// //n := hits["/doc/"]["au"]

	// // This approach becomes unwieldly when adding data, as for any given outer key you must check if the inner map exists,
	// // and create it if needed:

	// add(hits, "/doc/", "au")
	// add(hits, "/doc/", "au")
	// fmt.Println(hits)
	// fmt.Println(hits["/doc/"])
	// fmt.Println(hits["/doc/"]["au"])

		// however the dsign that uses a single map with a does away all that commplexity
		
	
}

func add(m map[string]map[string]int, path, country string) {
	mm, ok := m[path]
	if !ok {
		mm = make(map[string]int)
		m[path] = mm
	}
	mm[country]++
}

