// Problem 1: File Operations and Error Handling
// Task: Write a program that reads a text file containing a list of names (one per line),
// sorts the names alphabetically, and writes the sorted list to a new file.
// Handle all potential errors gracefully.

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <filepath>")
		return
	}

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file %v\n %v", filename, err)
		return
	}
	defer file.Close()

	filescan := bufio.NewScanner(file)
	filescan.Split(bufio.ScanLines)
	data := []string{}

	for filescan.Scan() {
		data = append(data, filescan.Text())
	}

	// check for the error that occurred during the scanning.
	if err := filescan.Err(); err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	sort.Strings(data)

	outputfile := "clean.txt"
	err = os.WriteFile(outputfile, []byte(fmt.Sprintf("%s\n", data)), 0o777)
	if err != nil {
		fmt.Println("error writing file")
		return
	}
	fmt.Printf("Sorted names written to %s\n", outputfile)
}

// func SortString(input []string) []string {
// 	n := len(input)

// 	if n < 1 {
// 		return []string{}
// 	}

// 	for i := 0; 1 < n-1; i++ {
// 		for j := 0; j < n-i-1; j++ {
// 			if input[i] > input[j+1] {
// 				input[j], input[j+1] = input[j+1], input[j]
// 			}
// 		}
// 	}
// 	return input
// }
