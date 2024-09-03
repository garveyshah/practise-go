package main

// A method is just a function with a special receiver type btwn the func and the method name.
// The reiver can either be a struct type or non-struct type.

// Syntax of a method declaration:
// func (t Type) methodName(parameter list) {
// 	//
// }

// The above snippet creates a method named methodName with receiver type Type. t is called as
//  the receiver and it can be accessed within the method.

import (
	"fmt"
)

type Employee struct {
	name     string
	salary   int
	currency string
}

// displaySalary() method has Emplayee as the receiver type

func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}

func main() {
	empl := Employee{
		name:     "Sam Adolf",
		salary:   5000,
		currency: "$",
	}
	empl.displaySalary()

}
