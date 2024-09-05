package main

import "fmt"

type Creature struct {
	Name     string
	Greeting string
}

func (c Creature) Greet() Creature {
	fmt.Printf("%s says %s!\n", c.Name, c.Greeting)
	return c
}

func (c Creature) SayGoodbye(name string) {
	fmt.Println("Farewell", name, "!")
}

func main() {
	sammy := Creature{
		Name:     "Sammy",
		Greeting: "Hello!",
	}
	sammy.Greet().SayGoodbye("gophers")

	Creature.SayGoodbye(Creature.Greet(sammy), "gophers")
}
