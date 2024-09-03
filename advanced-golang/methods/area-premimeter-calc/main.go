package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Shape interfaces
type Perimeter interface {
	Perimeter() float64
}

type Area interface {
	Area() float64
}

// Structs
type Circle struct {
	radius float64
}

type Rectangle struct {
	length float64
	width  float64
}

type Triangle struct {
	hypotenuse float64
	height     float64
	base       float64
}

// Area Methods
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// Perimeter methods
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (t Triangle) Perimeter() float64 {
	return t.base + t.height + t.hypotenuse
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

// Get Shape Data
func getData(shape string) (interface{}, error) {
	reader := bufio.NewReader(os.Stdin)

	switch strings.ToLower(shape) {
	case "rectangle":
		fmt.Println("Enter the dimensions of the rectangle in this order: <length> <width>")
		return getRectangleData(reader)
	case "triangle":
		fmt.Println("Enter the dimensions of the triangle in this order: <base> <height> <hypotenuse>")
		return getTriangleData(reader)
	case "circle":
		fmt.Println("Enter the radius of the circle:")
		return getCircleData(reader)
	default:
		return nil, fmt.Errorf("unknown shape type")
	}
}

func getRectangleData(reader *bufio.Reader) (Rectangle, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return Rectangle{}, err
	}
	inputs := strings.Fields(input)
	if len(inputs) != 2 {
		return Rectangle{}, fmt.Errorf("invalid input format\nUsage: <length> <width>")
	}
	length, err1 := strconv.ParseFloat(inputs[0], 64)
	width, err2 := strconv.ParseFloat(inputs[1], 64)
	if err1 != nil || err2 != nil {
		return Rectangle{}, fmt.Errorf("invalid number format")
	}
	return Rectangle{length: length, width: width}, nil
}

func getTriangleData(reader *bufio.Reader) (Triangle, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return Triangle{}, err
	}

	inputs := strings.Fields(input)
	if len(inputs) != 3 {
		return Triangle{}, fmt.Errorf("invalid input formart\nUsage: <base> <height> <hypotenuse>")
	}

	base, err1 := strconv.ParseFloat(inputs[0], 64)
	height, err2 := strconv.ParseFloat(inputs[1], 64)
	hypotenuse, err3 := strconv.ParseFloat(inputs[2], 64)
	if err1 != nil || err2 != nil || err3 != nil {
		return Triangle{}, fmt.Errorf("invalid number format")
	}
	return Triangle{hypotenuse: hypotenuse, base: base, height: height}, nil
}

func getCircleData(reader *bufio.Reader) (Circle, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("invalid input formart\nUsage: <radius>")
		return Circle{}, err
	}

	inputs := strings.Fields(input)
	radius, err1 := strconv.ParseFloat(inputs[0], 64)
	if err1 != nil {
		return Circle{}, fmt.Errorf("invalid number format")
	}
	return Circle{radius: radius}, nil
}

func main() {
	fmt.Println("Enter operation and shape (e.g., 'area circle'): ")
	reader := bufio.NewReader(os.Stdin)
	request, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	parts := strings.Fields(strings.TrimSpace(request))
	if len(parts) < 2 {
		fmt.Println("Usage: <operation> <shape>")
		return
	}

	operation := strings.ToLower(parts[0])
	shape := strings.ToLower(parts[1])

	shapeData, err := getData(shape)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	switch operation {
	case "area":
		if s, ok := shapeData.(Area); ok {
			fmt.Printf("The perimeter of the %s is %.2f\n", shape, s.Area())
		} else {
			fmt.Println("Error: Shape does not support are calculation")
		}
	case "perimeter":
		if s, ok := shapeData.(Perimeter); ok {
			fmt.Printf("The perimeter of the %s is %.2f\n", shape, s.Perimeter())
		} else {
			fmt.Println("Error: Shape does not suppurt perimeter calculation")
		}
	default:
		fmt.Println("Unkown operation")
	}
}
