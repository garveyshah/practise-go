package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter number!")

	reader := bufio.NewReader(os.Stdin)
	num, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input!!")
		return
	}

	num1 := strings.TrimSuffix(num, "\n")

	num2, err := strconv.Atoi(num1)
	if err != nil {
		fmt.Println("Input Error!! :- Please input a number!")
	}

	fmt.Println(GenarateFizzBuzz(num2))
}

func GenarateFizzBuzz(num int) string {

	result := []string{}
	for i := 1; i <= num; i++ {
		if i%5 == 0 && i%3 == 0 {
			result = append(result, "FizzBuzz ")
		} else if i%5 == 0 {
			result = append(result, "Buzz ")
		} else if i%3 == 0 {
			result = append(result, "Fizz ")
		} else {
			num2 :=strconv.Itoa(i)
			result = append(result, num2+" " )
		}
	}
	return strings.Join(result,"")
}
