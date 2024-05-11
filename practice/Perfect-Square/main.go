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
		log.Println("Error!! : ", err)
	}
}

func main() {
	fmt.Println("Enter number :...")

	reader := bufio.NewReader(os.Stdin)
	num, err := reader.ReadString('\n')
	CheckErr(err)

	num = strings.TrimSuffix(num, "\n")

	num1, err := strconv.Atoi(num)
	CheckErr(err)

	result := PerfectSquare(num1)

	fmt.Printf("%d is %s.\n", num1, result)
}

func PerfectSquare(num int) string {
	if num < 1 {
		return "Please enter positive number"
	}

	result := ""
	for i := 1; i <= num; i++ {
		if i*i == num {
			result = "a Perfect Square"
			break
		} else {
		result = "not a Perfect Square"
	}
}
	return result
}
