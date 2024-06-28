/* 
	Day 5: Implement a function to count the number of vowels in a string.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Input the string of characters.")

	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error!! reading input: ", err)
		return
	}

	str = strings.TrimSpace(str)
	fmt.Println()
	fmt.Println(str)

	fmt.Printf("\nVowel count is := %d\n", VowelCounter(str))
}

// funcion Vowel counter counts the vowels in a string
func VowelCounter(str string) int {
	vowelMap := map[rune]rune{'a': 'a', 'e': 'e', 'i': 'i', 'o': 'o', 'u': 'u'}
	str1 := strings.ToLower(str)
	count := 0
	for _, char := range str1 {
		if _, ok := vowelMap[char]; ok {
			count++
		}
	}
	return count
}
