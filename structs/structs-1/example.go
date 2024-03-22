package main

import "fmt"

type student struct {
	name string 
	age int
}

func ChangeName(pointer *student, nameChosen string) {

		pointer.name = nameChosen
}

func ChangeAge(pointer *student, ageChosen int) {

	pointer.age = ageChosen
}
func main () {
	chris := student{}

	fmt.Println(chris)

	chris.name = "chris"

	chris.age = 30

	fmt.Println(chris)

	ChangeName(&chris, "Augusto")

	fmt.Println(chris)

	ChangeAge(&chris, 50)

	fmt.Println(chris)

}