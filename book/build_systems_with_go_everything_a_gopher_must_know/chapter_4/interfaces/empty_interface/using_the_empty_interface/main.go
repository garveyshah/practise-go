// The epmty interface is a very ambiguous context for a typed language like Go.

package main

import "fmt"

func main() {
	var aux interface{} // intialize empty interface type

	fmt.Println(aux)

	aux = 10 // give it a value
	fmt.Println(aux)

	aux = "hello" // update the value
	fmt.Println(aux)
}

