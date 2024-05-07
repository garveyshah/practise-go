package main

import (
	"bufio"
	"fmt"
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
		// If an error occurs while reading,
		fmt.Println("")
		return 
	}

	// if len(os.Args) < 2 {
	// 	log.Fatal("Input less than 2")
	// }
	// args := os.Args[1]
	// args = strings.TrimSuffix(args, "\n")

	nums = strings.TrimSuffix(nums, "\n")

	numbers := strings.Fields(nums)
	if len(numbers) > 1 {
		fmt.Println("Error: Please insert only number")
		return
	}

	num, err := strconv.Atoi(nums)
	if err != nil {
		fmt.Println("Error: Please insert a number")
		return
	}

	if num < 0 {
		fmt.Println("Error!! :- Please enter a postive number.")
	}

	fmt.Printf("Sum of Ints from 1 upto %d is %d.\n", num, AddUptoNthTerm(num))
}

func AddUptoNthTerm(num int) int {
	result := 0

	for i := 1; i <= num; i++ {
		result += i
	}
	return result
}
