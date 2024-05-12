package main

import (
	"os"
	"io"
	"fmt"
)

func main() {
	// Check if the correct number of command-line arguments is provided.
	if len(os.Args) != 2 {
		fmt.Println("Please provide a filename.")
		return
	}

	// Extract the filename from the command-line arguments.
	filename := os.Args[1]

	// Open the file for reading.
	f, err := os.Open(filename) 
	// If there's an error in opening the file print error and exit.
	if err != nil {
		fmt.Printf("error opening %s: %s", filename, err) 
		return
	}
	// Ensure the file is closed when the function returns.
	defer f.Close()

	// Copy the content of the file to os.Stdout (terminal).
	_, err = io.Copy(os.Stdout, f)
	if err != nil {
		// If there's an error copying the file content,  print error and exit.
		fmt.Printf("error reading file: %s", err)
		return
	}

}