package main

import "fmt"

func main() {
	Plate(chicken, "rice")
	Plate(dal, "rice")
}

func Plate(f func() string, primaryDish string) {
	fmt.Println("This Plate Will Hve - " + primaryDish + f())
}

func chicken() string {
	return " with Chicken" // This plate Will Have - rice with Chicken
}

func dal() string {
	return " with Dal" // This Plate Will Have - rice with Dal
}
