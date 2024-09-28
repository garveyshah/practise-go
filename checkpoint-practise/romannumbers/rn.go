package main

import (
	"fmt"
	"os"

	"project01/checkpoint-practise/atoi/atoi"
)

// Convert an integer to a ROman numeral with step tracking
func IntToroman(num int) (string, string, error) {
	if num <= 0 || num >= 4000 {
		return "", "", fmt.Errorf("cannot convert to roman digit")
	}

	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	calculation := []string{"M", "M-C", "D", "D-C", "C", "C-X", "L", "L-X", "X", "X-I", "V", "V-I", "I"}
	number := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	
	roman := ""
	calc := ""

	for i := 0; i < len(number); i++ {
		for num >= number[i] {
			roman += symbols[i]
			if len(calculation[i]) > 1 {
				calc += "(" + calculation[i] + ")" + "+"
			} else {
				calc += calculation[i] + "+"
			}
			num -= number[i]
			 //num = 2000 - 1000
		}
	}

	if calc[len(calc)-1] == '+' {
		calc = calc[:len(calc)-1]
	}
	return roman, calc, nil
}

func main() {
	if len(os.Args) < 2 {
		return
	}

	num, err := atoi.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	roman, steps, err := IntToroman(num)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(roman)
	fmt.Println(steps)
}
