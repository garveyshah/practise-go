package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Abs written as a method.
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	r := Vertex{5, 6}
	fmt.Println(Abs(r))

}

// Abs Written as a regular function.
func Abs( r Vertex) float64 {
	return math.Sqrt(r.X *r.X + r.Y*r.Y)
}
