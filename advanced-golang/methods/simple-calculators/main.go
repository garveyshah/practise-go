package main

import (
	"fmt"
	"os"
	"strconv"
)

type Data struct {
	Sign    string
	number1 int
	number2 int
}

func (d Data) calculator() {
	var result int
	switch d.Sign {
	case "-":
		result = d.number1 - d.number2
		fmt.Printf("%d - %d = %d\n", d.number1, d.number2, result)
	case "+":
		result = d.number1 + d.number2
		fmt.Printf("%d - %d = %d\n", d.number1, d.number2, result)
	case "*":
		result = d.number1 + d.number2
		fmt.Printf("%d * %d = %d\n", d.number1, d.number2, result)
	case "/":
		result = d.number1 / d.number2
		fmt.Printf("%d/%d = %d\n", d.number1, d.number2, result)
	default:
		fmt.Println("Unsupported operation")
		return
	}
}

func isEmpty(d Data) bool {
	return d.Sign == "" && d.number1 == 0 && d.number2 == 0
}

func (d *Data) getData() {
	if len(os.Args) != 4 {
		fmt.Println("Usage : go run . <number 1> <operator> <number2>\ngo run . 1 + 2")
		return
	}
	// }
	var err error
	d.Sign = os.Args[2]
	d.number1, err = strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("error invalid character %v", string(os.Args[1]))
		return
	}
	d.number2, err = strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Printf("error invalid character %v", string(os.Args[1]))
		return
	}

}

func main() {
	data := Data{}
	data.getData()
	if isEmpty(data) {
		return
	}
	data.calculator()
}
