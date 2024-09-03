package main

import (
	"fmt"
	"math"
)

// A method is just a function with a special receiver type btwn the func and the method name.
// The reiver can either be a struct type or non-struct type.

// Syntax of a method declaration:
// func (t Type) methodName(parameter list) {
// 	//
// }

// The above snippet creates a method named methodName with receiver type Type. t is called as
//  the receiver and it can be accessed within the method.

// type Employee struct {
// 	name     string
// 	salary   int
// 	currency string
// }

// // displaySalary() method has Emplayee as the receiver type

// func (e Employee) displaySalary() {
// 	fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary)
// }

// func main() {
// 	empl := Employee{
// 		name:     "Sam Adolf",
// 		salary:   5000,
// 		currency: "$",
// 	}
// 	empl.displaySalary()

// }

// AREA

type Rectangle struct {
	length int
	width  int
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	r := Rectangle{
		length: 10,
		width:  5,
	}
	fmt.Printf("Area of rectangle %d\n", r.Area())
	c := Circle{
		radius: 12,
	}
	fmt.Printf("Area of circle %f\n", c.Area())
}
