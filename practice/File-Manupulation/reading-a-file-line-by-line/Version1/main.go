package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

// func LineByLine reads a file line line and prints its contents to stdout(command-line)
func LineByLine(file string) error {
	// Open the file
	fd, err := os.Open(file)
	if err != nil {
		// If there's an error openig the file, return it with context.
		return fmt.Errorf("failed to open file: %w", err)
	}
	//Ensure the file is closed after the function exits.
	defer fd.Close()
	// Create a newbuffered reader to read from the file.
	reader := bufio.NewReader(fd)
	for {
		// Read each line from the file.
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			// If end of the  file is reached break the loop.
			break
		} else if err != nil {
			// If there's an error reading form the file, print the error.
			fmt.Printf("error reading file %s", err)
			break
		}
		// Print the line to standard output(command-line)
		fmt.Print(line)
	}
	return nil
}

func main() {
	// Parse command-line arguments.
	flag.Parse()
	// Check if no arguments are provided.
	if len(flag.Args()) == 0 {
		// Print usage information if no arguments are provided.
		fmt.Printf("usage: lByL text1.txt[file1]\n")
		return
	}
	// Iterate over the provided file names.
	for _, file := range flag.Args() {
		// Call LineByLine function for each file.
		err := LineByLine(file)
		if err != nil {
			// If there's an error processing the file, print the error
			fmt.Println(err)
		}
	}
}
