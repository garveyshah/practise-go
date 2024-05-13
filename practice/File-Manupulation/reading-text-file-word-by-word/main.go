package main

import (
	"fmt"
	"os"
	"bufio"
	"flag"
	"io"
	"regexp"
)

func WordByWord(file string) error {
	// Open the specifed file Reading.
	fd, err := os.Open(file)
	// If there is an error Opening the file print the error and exit.
	if err != nil {
		return err
	}
	
	// Ensure the file is closed when the function returns
	defer fd.Close()

	// Create a reader to read the file save in (fd).
	reader :=  bufio.NewReader(fd)
	for {
		// Read the file line by line. 
		line, err := reader.ReadString('\n')
		// If there is an error reading the file Break.
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			return err
		}
		r := regexp.MustCompile(`[^\s]+`)
	words := r.FindAllString(line, -1)
	for i := 0; i < len(words); i++ {
		fmt.Println(words[i])
	}
	
	}
	return nil
}

func main() {
	// Parse the arguments on the command line.
	flag.Parse()
	// Check if the required number of arguments are given if not print usage instructions and exit.
	if len(flag.Args()) == 0 {
		fmt.Printf("usage: main <file1> [<file2> ...]\n")
		return
	}
	// Iterate over the provided number of arguments (file names)
	for _, file := range flag.Args() {
		// Call WorfByWord function on each of the file names.
		err := WordByWord(file)
		if err != nil {
			fmt.Println(err)
		}
	}
}