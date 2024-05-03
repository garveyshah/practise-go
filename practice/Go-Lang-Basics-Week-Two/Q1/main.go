//Write a Go program that prompts the user to enter a number 
//and then prints whether it is even or odd.

package main

import "fmt"

func main() {
	num := 9
	num2 := 4	
	fmt.Println(IsOddOrEven(num))
	fmt.Println(IsOddOrEven(num2))

}

func IsOddOrEven(num int) string {
		if num%2 == 0 {
			
			return "Number is Even"
		} 
		return "Number is Odd"
	
}