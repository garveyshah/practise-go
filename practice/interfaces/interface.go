package main

import "math"

type square struct {
	length float64
}

type circle struct {
	radius float64
}

// square methods
func (s square) area() float64 {
	return s.length * s.length
}

func (s square) circumf() float64 {
	return s.length * 4
}

// cirlce methods
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius 
}

func (c circle) circumf() flaot64 {
	return 2 * math.Pi * c.radius
}

fucn printShapeInfo(s) {
	fmt.Printf("area of %T: %0.2f\n", s, s.area)
	
}