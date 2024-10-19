package main

import "fmt"

// func main() {
// 	fmt.Println("Structs in golang")
// 	// no inheritance in golang; NO super or parent

// 	godwin := User{"Godwin", "godwin@go.dev", true, 24}
// 	fmt.Println(godwin)
// 	fmt.Printf("godwin details are: %+v\n", godwin)
// 	fmt.Printf("Name is %v and email is %v.", godwin.Name, godwin.Email)
// }

type User struct {
	Name   string
	Email  string
 	Status bool
	Age    int
 }
func main() {
	var p1 Point = Point{1,2}
	var p2 Point = Point{-5,7}
	fmt.Println(p1.x, p1.y)
	fmt.Println(p2.x, p2.y)


}

	type Point struct {
		x int32
		y int32
	}

	