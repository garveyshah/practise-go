package main

import (
	"fmt"
	"interfaces/donald"
	"interfaces/daisy"
)

func MakeDuckQuack(d donald.Duck) {
	fmt.Println(d.Quack())
}

type DonaldDuck struct{}

func (DonaldDuck) Quack() string {
	return "Quack from Donald!"
}

type DaisyDuck struct {}

func (DaisyDuck) Quack() string {
	return "Quack from Daisy!"
}

func main() {
	var donaldDuck donald.Duck = DonaldDuck{}
	MakeDuckQuack(donaldDuck)
	var daisyDuck daisy.Duck = DaisyDuck{}
	MakeDuckQuack(daisyDuck)
}
 