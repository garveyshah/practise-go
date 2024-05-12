package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please give one argument.")
		return
	}

	path := args[1]

	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println("Path does not exist!", err)
	}

	mode := fileInfo.Mode()
	if mode.IsRegular() {
		fmt.Println(path, "is a regular file!")
	}
}