package main

import (
	"fmt"
	"project01/100daycodechallange/day19/calc"
)

func main() {
	// if len(os.Args) != 3 {
	// 	fmt.Println("Usage : go run . <number>\nExample : go run . 90")
	// 	return
	// }

	// num, err := strconv.Atoi(os.Args[1])
	// if err != nil {
	// 	fmt.Println("Err : ", err)
	// 	return
	// }

	// exp, err := strconv.Atoi(os.Args[2])
	// if err != nil {
	// 	fmt.Println("Err : ", err)
	// 	return
	// }

	// power := calc.Power(num, exp)

	// fmt.Printf("%d raised to the power %d is \"%d\"\n", num, exp, power)
	fmt.Println(calc.Factors(180))
}
