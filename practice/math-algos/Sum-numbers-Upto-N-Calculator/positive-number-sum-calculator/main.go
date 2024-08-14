package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Promt the user to enter a number (N).
	fmt.Println("Hello Please Enter Number(N)")

	// Create a reader to read input from standart input(keyboard).
	reader := bufio.NewReader(os.Stdin)

	// Read input until a newline character ('\n') is encountered.
	nums, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("")
		return
	}

	// Trim newline character from the output.
	nums = strings.TrimSuffix(nums, "\n")

	// Trim input string by spaces to check if it contains more than number.
	numbers := strings.Fields(nums)
	if len(numbers) > 1 {
		fmt.Println("Error: Please insert only number")
		return
	}

	// Convert the input string to an integer.
	num, err := strconv.Atoi(nums)
	if err != nil {
		fmt.Println("Error: Please insert a number")
		return
	}

	// Validate if the output is a positive integer.
	if num < 0 {
		fmt.Println("Error!! :- Please enter a postive number.")
	}

	// Calcutlate and print the sum of Positive integers from 1 upto the input number.
	fmt.Printf("Sum of Ints from 1 upto %d is %d.\n", num, SumOfPositiveInts(num))
}

// Calcutlate and print the sum of Positive integers from 1 upto the input number.
func SumOfPositiveInts(num int) int {
	result := 0

	// Check if num is at the maximum value for the respective integer type.
	if num == math.MaxInt32 || num == math.MaxInt64 {
		// Calculate the sum using a formula to avoid overflow.
		return num * (num + 1) / 2
	}

	for i := 1; i <= num; i++ {
		result += i
	}
	return result
}
