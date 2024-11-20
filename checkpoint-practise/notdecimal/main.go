
/*
 notdecimal
Instructions

Write a function called NotDecimal() that takes as an argument a string in
 form of a float number with the decimal point and returns a string converted
  to int without the decimal point (you will have to multiply it by 10^n to remove the .).

    If the number doesn't have a decimal point or there is only a zero after the .
	 	return the number followed by a newline \n.
    If the argument is empty return a newline \n.
    If the argument is not a number return it followed by a newline \n.

Expected function

func NotDecimal(dec string) string {

}
 */
package main

import "fmt"

func main() {
	fmt.Print(NotDecimal("0.1"))
	fmt.Print(NotDecimal("174.2"))
	fmt.Print(NotDecimal("0.1255"))
	fmt.Print(NotDecimal("1.20525856"))
	fmt.Print(NotDecimal("-0.0f00d00"))
	fmt.Print(NotDecimal(""))
	fmt.Print(NotDecimal("-19.525856"))
	fmt.Print(NotDecimal("1952"))
	// fmt.Print(NotDecimal("2"))
}


func NotDecimal(dec string) string {
	// return a \n if the argument is empty
	if dec == "" {
		return "\n"
	}

	// if the number is a single digit, or charater return the number'
}




























// func NotDecimal(s string) (result string) {
// 	if s == "" {
// 		return "\n"
// 	}

// 	if len(s) == 1 {
// 		return s
// 	}

// 	for _, char := range s {
// 		if(char < '0' || char > '9') {
// 			 if char != '-' && char != '.' {
// 			return s + "\n......"
// 			 }
// 		}
// 		if char != '.' {
// 			result += string(char)
// 		}
// 	}

// 	for result[0] == '0' && len(result) > 1 {
// 		result = result[1:]
// 	}

// 	if result == "0" {
// 		return "\n"
// 	}
// 	return result +  "\n"
// }

