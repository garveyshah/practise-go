package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Check if the correct number of command-line arguments is provided.
	if len(os.Args) != 2 {
		fmt.Println("Please provide a filename")
		return
	}
	// Extract the filename from the command-line arguments.
	filename := os.Args[1]

	// Open the file by the specified filename for reading and stores in f .
	f, err := os.Open(filename)
	fmt.Println(f)
	if err != nil {
		// If there's an error opening the file, print the error and exit.
		fmt.Printf("error opening %s: %s", filename, err)
		return
	}
	//  Ensure the file is closed when then function returns.
	defer f.Close()

	// make a byte slice with a specifed size to store the read content.
	buf := make([]byte, 450)

	// Read exactly the specifed of content from the file(`f`) into the buffer(`buf`).
	_, err = io.ReadFull(f, buf)
	// if there is and errr
	if err != nil {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
		fmt.Printf("Error reading file: %s\n", err)
		return
	}
	fmt.Println(string(buf))
}
