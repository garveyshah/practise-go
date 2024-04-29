package main

import "fmt"

func AddTen(nbr int) int {
	return nbr + 10
}

func AddTwenty(nbr int) int {
	return nbr + 20
}

func AddThirty(nbr int) int {
	return nbr + 30
}



func main() {

	result := 0
	result = AddTwenty(result)
	fmt.Println(result)

}
