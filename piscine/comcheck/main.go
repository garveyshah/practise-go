package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	for _, char := range args {
		if char == "01" || char == "galaxy" || char == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}
