package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func lineByLine(file string) error {
	// Open the file for reading
	fd, err := os.Open(file)
	if err != nil {
		// If there's an error opening the file, return it with context
		return fmt.Errorf("failed to open file: %w", err)
	}
	// Ensure the file is closed after the function exits
	defer fd.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		// Print each line of the file
		fmt.Println(scanner.Text())
	}
	// Check for any errors encountered during scanning
	if err := scanner.Err(); err != nil {
		// If there's an error scanning the file, return it with context
		return fmt.Errorf("error reading file: %w", err)
	}
	// If everything executed successfully, return nil
	return nil
}

func main() {
	// Parse command-line arguments
	flag.Parse()
	// Check if there are no arguments provided
	if flag.NArg() == 0 {
		// Print usage information if no arguments are provided
		fmt.Printf("usage: lByL text1.txt [file1]\n")
		return
	}
	// Iterate over the provided file names
	for _, file := range flag.Args() {
		// Call lineByLine function for each file
		if err := lineByLine(file); err != nil {
			// If there's an error processing the file, print it
			fmt.Println(err)
		}
	}
}
