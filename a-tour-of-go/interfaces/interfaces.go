package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type myFloat float64

func (f myFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var abser Abser

	//f := myFloat(-math.Sqrt2)

	v := Vertex{3, 4}

	//abser = f
	abser = &v

	// abser wich implements the method Abs has been give f - object of type float64 as a value

	fmt.Println(abser.Abs())
}


type writer interface {
	Write(p []byte) (n int, err error)
}

type Stringer interface {

}

https://go.dev/tour/methods/9 