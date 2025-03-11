package main

import "fmt"

type Greeter interface {
	SayHello() string
}

type Person struct {
	name string
}

func (p *Person) SayHello() string {
	return "Hi! This is me " + p.name
}

// You may consider implementing an interface using methods with pointers and value receivers
// simultaneously to be able to work with both flavors.
// Something like
// func (p Person) SayHello() string {
// 	return "Hi! This is me" + p.name
// }
func main() {
	var g Greeter

	p := Person{"John"}

	// g = p // -> Does not work
	g = &p

	fmt.Println(g.SayHello())
}


