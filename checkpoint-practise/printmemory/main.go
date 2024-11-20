package main

import (
	"fmt"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := [10]byte{}

	for j := 0; j < 5; j++ {
		for i := 0; i < random.IntBetween(7, 10); i++ {
			table[i] = byte(random.IntBetween(13, 126))
		}
		challenge.Function("PrintMemory", PrintMemory, solutions.PrintMemory, table)
	}
	table2 := [10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'}
	challenge.Function("PrintMemory", PrintMemory, solutions.PrintMemory, table2)

	res := PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
	fmt.Println(res)
}

// func PrintMemory(arr [10]byte) string {
// 	res := ""
// 	// Print the hexadecimal representation
// 	for i, b := range arr {
// 		res += printHex(b)
// 		res += string(' ')
// 		if (i == 3) || (i == 7) || (i == 9) {
// 			res += "\n"
// 		}
// 	}

// 	// Print the ASCII representation
// 	for _, b := range arr {
// 		if b >= 32 && b <= 126 {
// 			res += string(rune(b))
// 		} else {
// 			res += "."
// 		}
// 	}
// 	// res += "\n"
// 	return res
// }

// func printHex(b byte) string {
// 	hex := ""
// 	hexDigits := "0123456789abcdef"

// 	hex += string(rune(hexDigits[b>>4]))
// 	hex += string(rune(hexDigits[b&0x0f]))
// 	// z01.PrintRune(' ')
// 	return hex
// }
