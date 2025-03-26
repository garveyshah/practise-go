package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("The quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}

func CountDuplications(s string, i int) int {
	var counter int 

	for _, r := range s[i:] {
		if r == rune(s[i]) {
			counter++
		}
		break
	}
	return counter
}

func ZipString(s string) string {
	var result string 

	i := 0 
	for i < len(s) {
		counter := CountDuplications(s, i)
		result = result + strconv.Itoa(counter) + string(s[i])
		i += int(counter)
	}
	return result
}