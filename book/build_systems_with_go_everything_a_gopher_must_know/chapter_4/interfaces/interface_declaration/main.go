/*  
The example below declares the `Animal interface` with two methods `Roar` and `Run`.
Next we have `dog` and `cat` types that defines these methods.
Automatically both types are considered to implement interface `Animal`.
Function `RoarAndRun` receives an `Animal` type, so we invoke the `Roar` and `Run` methods independently of the final 
argument type.

Notice that method receivers from 'Dog' and 'Cat' are different.
Because all the methods of the 'interface' must be implemented in order to be considered a type to implement
'interface' of type 'Animal', certain combinations in the example can faile

'RoarAndRun(myCat)' does not work because the receivers of the methods in the 'cat' type are pointers while we pass an argument by value.
In case the method assigns a new value to any field, this cannot be reflected in the origina; caller because its a copy. 

*/

package main

import "fmt"

type Animal interface {
	Roar() string
	Run() string
}

type Dog struct {}

func (d Dog) Roar() string {
	return "Woof!"
}

func (d Dog) Run() string {
	return "run like a dog"
}

type Cat struct {}

func (c *Cat) Roar() string {
	return "meow!"
}

func (c *Cat) Run() string {
	return "run like a cat"
}

func RoarAndRun(a Animal) {
	fmt.Printf("%s and %s\n", a.Roar(), a.Run())
}

func main() {
	myDog := Dog{}
	myCat := Cat{}

	RoarAndRun(myDog)
	RoarAndRun(&myCat)
}