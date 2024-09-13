package main

import "fmt"

func main() {
	res := PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
	fmt.Println(res)
}

func PrintMemory(arr [10]byte) string {
	res := ""
	// Print the hexadecimal representation
	for i, b := range arr {
		res += printHex(b)
		res += string(' ')
		if (i == 3) || (i == 7) || (i == 9) {
			res += "\n"
		}
	}

	// Print the ASCII representation
	for _, b := range arr {
		if b >= 32 && b <= 126 {
			res += string(rune(b))
		} else {
			res += "."
		}
	}
	//res += "\n"
	return res
}

func printHex(b byte) string {
	hex := ""
	hexDigits := "0123456789abcdef"

	hex += string(rune(hexDigits[b>>4]))
	hex += string(rune(hexDigits[b&0x0f]))
	//z01.PrintRune(' ')
	return hex
}
