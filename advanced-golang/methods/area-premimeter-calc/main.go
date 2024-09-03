package main

import (
	"math"
)

type Circle struct {
	radius float64
}

type Rectangle struct {
	length int
	width  int
}

type Triangle struct {
	height int
	base   int
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (t Triangle) Area() int {
	return t.base / 2 * t.height
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func (c Circle) Perimeter() float64 {
	return math.Pi * 2 * c.radius
}

func (t Triangle) Perimeter() int {
	return t.base + t.height
}