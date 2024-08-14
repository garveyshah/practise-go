package main

import "fmt"

func main() {
	// Function as Parameters in Golang
	Plate(chicken, "rice")
	Plate(dal, "rice")

	// Function as a return value in Golang

	// het function as return value
	fn := plate("rice")

	// call the return value as the function
	fn(" with chicken")
	fn(" with fish")
}

// Function as Parameters in Golang

func Plate(f func() string, primaryDish string) {
	fmt.Println("This Plate Will Hve - " + primaryDish + f())
}

func chicken() string {
	return " with Chicken" // This plate Will Have - rice with Chicken
}

func dal() string {
	return " with Dal" // This Plate Will Have - rice with Dal
}

// Function as a return value in Golang

func plate(primaryDish string) func(param string) {
	// function is treated as a return value
	return func(param string) {
		fmt.Println("This plate Will have - " + primaryDish + param)
	}
}
