package main

import (
	"errors"
	"fmt"
)

type Rectangle struct {
	Height int
	Width  int
}

func NewRectangle(height int, width int) (*Rectangle, error) {
	if height <= 0 || width <= 0 {
		return nil, errors.New("params must be greater than zero")
	}
	return &Rectangle{height, width}, nil
}

func main() {
	a := Rectangle{Height: 7}
	fmt.Println(a)

	r, err := NewRectangle(2, 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
	fmt.Println(*r)
}
