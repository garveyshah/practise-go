package main

import (
	"bufio"
	"fmt"
	"os"
	"project01/100daycodechallange/day1/evenorodd"
	"strings"
)

func main() {
	fmt.Println("Enter a number:...")

	reader := bufio.NewReader(os.Stdin)
	num, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error!! reading input...", err)
		return
	}

	num = strings.TrimSpace(num)

	num1, err := customAtoi(num)
	if err != nil {
		fmt.Println("Error :-", err)
		return
	}

	fmt.Printf("\n%d is an \"%v\" number.\n",num1,(evenorodd.EvenOrOdd(num1)))
}

func customAtoi(s string) (int, error) {
	result := 0

	for _, char := range s {
		if char < '0' || char > '9' {
			return 0, fmt.Errorf("error invalid character :\"%v\"", char)
		}
		result = result*10 + int(char-'0')
	}
	return result, nil
}
