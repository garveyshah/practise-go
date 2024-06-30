/*
Day 7: Implement a function to check if two strings are anagrams.
*/


//workflow
//1. interate through the strings look for instances of "" if found append to separate string.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter the strings each enclosed in \"\" and separated by a space.")

	reader := bufio.NewReader(os.Stdin)
	MainStr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input", err)
		return
	}

	MainStr = strings.TrimSpace(MainStr)

	Str1, Str2 :=GetInput(MainStr)
	fmt.Printf("String 1 is :%v.\nString 2 is :%v.",Str1, Str2)



}

func GetInput(s string) (string, string){
	strSlice := strings.Split(s, " ")
	
	str1 := strings.Trim(strSlice[0],"")
	str2 := strings.Trim(strSlice[1],"")

	return str1, str2
}
