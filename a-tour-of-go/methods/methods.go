package main

import (
	"fmt"
	"math"
)

// A method is function with a special reciever argument.
// The reciever appears in its own argument list between the func keyword and the method name
// A method is just a function with a receiver argument.


type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
fmt.Println("Hello World")
}
