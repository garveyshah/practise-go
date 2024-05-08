package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Prompt the user input a number.
	fmt.Println("Please enter Number: ")

	// Initate a variable reader to read the input string.
	reader := bufio.NewReader(os.Stdin)
	nums, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error!!: Errror reading input")
		return
	}
	// Trim the suffix newline character ('\n).
	nums = strings.TrimSuffix(nums, "\n")

		// Check if there is more than one number in the input string.
	numbers := strings.Fields(nums)
	if len(numbers) > 1 {
		fmt.Println("Errror!!: Please input 1 number")
		return
	}

// Convert input string to floating-point number.
num := StrToFloat(nums)

// Check if the input is a whole number.
isInteger := IsInt(num)

if !isInteger {
	fmt.Println("Error: Please input a whole number")
	return
}

// Convert the float number to an integer.
intNum := int(num)

// Determine if the number is odd or even.
fmt.Printf("%d is %s\n", intNum, OddOrEven(intNum))
}

func OddOrEven(num int) string {
	response := ""
	if num%2 == 0 {
		response = "Even"
		return response
	}
	response = "Odd"
	return response
}

func IsInt(x float64) bool {
		return x == float64(int64(x))
}

func StrToFloat(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return 0.0
	}

	return f
}
