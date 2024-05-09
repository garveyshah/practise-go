package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func CheckErr(err error) {
	if err != nil {
		log.Println("Error!! :-", err)
	}
}

func main() {
	// Prompt the user to enter a number to use when calculating factorial.
	fmt.Println("Hello Please key in the number to compute...")

	// Initiate a reader to read the users input.
	reader := bufio.NewReader(os.Stdin)
	num, err := reader.ReadString('\n')
	CheckErr(err)

	num = strings.TrimSuffix(num, "\n")

	num2, err := strconv.Atoi(num)
	CheckErr(err)

	result := RecursiveFactorial(num2)

	fmt.Println(result)
}

func RecursiveFactorial(num int) int {
	if num < 0 || num > 20  {
		return 0
	}
	if num == 1 || num == 0 {
		return 1
	} 
	return num * RecursiveFactorial(num-1)
}
