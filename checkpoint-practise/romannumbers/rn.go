package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please parse an argument")
		return
	}

	result := 0

	for _, char := range os.Args[1] {
		digit := int(char - '0')
		if digit < 0 || digit > 9 {
			fmt.Println("ERROR: cannot convert to roman digit.")
			return
		}

		result = result*10 + digit
	}

	if result == 0 || result == 4000 {
		fmt.Println("ERROR: cannot convert to roman digit.")
		return
	}

 RomanNums := map[int]string{
        1:  "I", 2:  "II", 3:  "III", 4:  "IV", 5:  "V",
        6:  "VI", 7:  "VII", 8:  "VIII", 9:  "IX", 10: "X",
        11: "XI", 12: "XII", 13: "XIII", 14: "XIV", 15: "XV",
        16: "XVI", 17: "XVII", 18: "XVIII", 19: "XIX", 20: "XX",
        21: "XXI", 22: "XXII", 23: "XXIII", 24: "XXIV", 25: "XXV",
        26: "XXVI", 27: "XXVII", 28: "XXVIII", 29: "XXIX", 30: "XXX",
		50: "L", 100: "C", 500: "D", 1000: "M",
    }
	
	
}
