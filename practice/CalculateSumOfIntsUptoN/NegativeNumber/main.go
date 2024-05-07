package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Prompt the user user to enter the number(N)
	fmt.Println("Hello please enter a number(N).")

	// Create a reader to read input from standart input (keybord).
	reader := bufio.NewReader(os.Stdin)
	// Read input until a newline character('\n') is encountered.
	num, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error: Reading input.")
		return
	}

	// Trim newline character from the output.
	num = strings.TrimSuffix(num, "\n")

	// Trim input string by spaces to check if it contains more than number.
	numbers := strings.Fields(num)
	if len(numbers) > 1 {
		fmt.Println("Error! :- Please enter only 1 number.")
		return
	}

	// Convert the input string to interger.
	nums, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println("Error! :- Please numbers only.")
	}

	if nums > 0 {
		fmt.Println("Error!! :- Please enter a negative number.")
		return
	}

	fmt.Printf("Sum of Digits for %d to 0, is %d\n", nums, AddNegsUptoNthTerm(nums))
}

func AddNegsUptoNthTerm(nums int) int {
	result := 0

	for i := nums; i <= 0 ; i++ {
		result += i
	}
	return result
}
