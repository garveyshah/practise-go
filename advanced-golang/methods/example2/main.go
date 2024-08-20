package main

import "fmt"

type Rectangle struct {
	height, width float64
}

// Method to calculate the area of the rectang;e
func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func main() {
	// Create a Rectangle instance
	r := Rectangle{height: 5, width: 6}

	// Call the Area method on the r instance
	fmt.Println("Area of the rectangle:", r.Area())
}
