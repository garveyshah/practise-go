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
		log.Println("Erro!!", err)
	}
}

func main() {
	fmt.Println("Enter number to work with")

	reader := bufio.NewReader(os.Stdin)
	num, err := reader.ReadString('\n')
	CheckErr(err)

	num = strings.TrimSuffix(num, "\n")

	num1, err := strconv.Atoi(num)
	CheckErr(err)

	fmt.Println("Number entered is:-",num1)
	// result := LargetPrimeFact(num1)
}

func LargetPrimeFact(num int) int {
	if num < 1 {
		return 1
	}

	if num%
	return num
}