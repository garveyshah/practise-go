package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 2 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	} 
	z01.PrintRune(rune(len(os.Args)))
	
	z01.PrintRune('\n')
}
