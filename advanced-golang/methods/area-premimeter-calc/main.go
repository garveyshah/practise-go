package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// type Circle struct {
// 	radius float64
// }

// type Rectangle struct {
// 	length int
// 	width  int
// }

// type Triangle struct {
// 	hypoteneus int
// 	height     int
// 	base       int
// }

// // Area Methods
// func (c Circle) Area() float64 {
// 	return math.Pi * c.radius * c.radius
// }

// func (t Triangle) Area() int {
// 	return t.base / 2 * t.height
// }

// func (r Rectangle) Area() int {
// 	return r.length * r.width
// }

// // Perimeter methods

// func (c Circle) Perimeter() float64 {
// 	return math.Pi * 2 * c.radius
// }

// func (t Triangle) Perimeter() int {
// 	return t.base + t.height + t.hypoteneus
// }

// func (r Rectangle) Perimeter() int {
// 	return 2*r.length + 2*r.width
// }

// // Interfaces

// type Perimeter interface {
// 	Perimeter()
// }

// type Area interface {
// 	Area()
// }

// func ()getData(shape target interface{}) interface{} {
// 	var shape string
// 	if len(os.args) < 2 {

// 	}
// 	switch len(os.Args) {
// 		case len(os.Args) < 1 && len(os.Args) >
// 	}
// }

// takes poiner to struct

func main() {
	fmt.Println("Enter shape and formal\nExample : area circle")
	reader := bufio.NewReader(os.Stdin)
	request, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error ", err)
	}

	shape := strings.Split(request, " ")[0]
	formula := strings.Split(request, " ")[1]

	 
	//	shape = strings.ToLower(shape)
	switch strings.ToLower(shape) {
	case "rectangle":
		fmt.Printf("Enter the dimentions of the %v in this order:  <length><space><width><height>\n5 6 7\n", shape)
	case "triangle":
		fmt.Printf("Enter the dimentions of the %v in this order:  <base><space><height><space><hypoteneuse>\n5 6 7\n", shape)
	case "circle":
		fmt.Printf("Enter the radius of the %v.\n", shape)
	}

	reader1 := bufio.NewReader(os.Stdin)
	data, err := reader1.ReadString('\n')
	if err != nil {
		fmt.Println("Error ", err)
	}


	fmt.Printf(formula, data, shape, "\n")

	
}
