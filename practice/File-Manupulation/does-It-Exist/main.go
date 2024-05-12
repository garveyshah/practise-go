package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("PLease give one argument.")
		return
	}
	path := args[1]

	_, err := os.Stat(path)
	if err != nil {
		fmt.Println("Path does not exist!", err)
	}
	fmt.Println("Path Exists.")
}
