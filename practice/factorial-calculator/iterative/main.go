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
		log.Println("Error!!", err)
	}
}

func main() {
	fmt.Println("Input number")

	reader := bufio.NewReader(os.Stdin)
	num, err := reader.ReadString('\n')
	CheckErr(err)

	num = strings.TrimSuffix(num, "\n")

	num2, err := strconv.Atoi(num)
	CheckErr(err)

	result := IterativeFactorial(num2)

	fmt.Printf("!%d is %d.\n", num2, result)
}

func IterativeFactorial(num int) int {
	if num < 0 || num > 20 {
		return 0
	} else if num == 0 || num == 1 {
		return 1
	}

	result := 1

	for i := 1; i <= num; i++ {
		result = result * i
	}

	return result
}
