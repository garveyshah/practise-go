/*
	When a struct is embedded into other structs its methods are made availbale to the second one.
	This act as some sort of inheritance in Go.
	This acts as some sort of inheritance in Go.
	In the example below, the type `Box` embeds the type `Rectangle`.
	Observe how the method `Volume` can directly invoke the `Surface` method from `Rectangle` to compute the
	volume of the box.

	NOTE: embedded methods only work if there is no ambiguity in its defination.
	in the example shown below, Invoking `Hi` in `Greeter` is not possible because the compiler cannot determine
	which type `A` or `B` is the owner.
	This has to be solved by specifying the methof owner.
*/

package main

import "fmt"

type Rectangle struct {
	Height int
	Width  int
}

func (r Rectangle) Surface() int {
	return r.Height * r.Width
}

type Box struct {
	Rectangle
	depth int
}

func (b Box) Volume() int {
	return b.Surface() * b.depth
}

type A struct {}

func (a A) Hi() string {
	return "A says Hi"
}

type B struct {}
func (b B) Hi() string {
	return "B says Hi"
}

type Greeter struct {
	A
	B
}

func (g Greeter) Speak() string {
	// return g.Hi() // -> This method belongs to A or B?
	return g.A.Hi() +  " " +  g.B.Hi()
}
func main() {
	b := Box{Rectangle{3, 3}, 3}
	fmt.Printf("%v+v\n", b)
	fmt.Println("Volume", b.Volume())

	d := Greeter{}
	fmt.Println(d.Speak())
}
