package main

import "fmt"

type student struct {
	name string
	age  int
}

func changeName(pointer *student, nameChosen string) {
	pointer.name = nameChosen
}

func changeAge(pointer *student, ageChosen int) {
	pointer.age = ageChosen
}

func main() {
	chris := student{}
	fmt.Println(chris)

	chris.name = "chris"

	chris.age = 30
	fmt.Println(chris)

	changeName(&chris, "Godwin")
	fmt.Println(chris)

	changeAge(&chris, 25)
	fmt.Println(chris)

}
