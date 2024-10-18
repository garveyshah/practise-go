package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter the text sequence to check")

	reader := bufio.NewReader(os.Stdin)
	text1, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error!!: error reading string:- ", err)
		return
	}
	newText := ""
	text := strings.TrimSpace(text1)
	text = strings.ToLower(text)

	for _, char := range text {
		if (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') {
			newText = newText + string(char)
		} 
	}

	revText := ""

	for j := len(newText) - 1; j >= 0; j-- {
		revText += string(newText[j])
	}

	if revText == newText {
		fmt.Printf("\n%v IS A PALINDROME.\n", text1)
		return
	} else {
		fmt.Printf("\n%v IS NOT A PALINDROME.\n", text1)
		return
	}
}
