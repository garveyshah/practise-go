package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello Medium")
	OutPut()
}

func OutPut() {
	// Empty buffer (implements io.Writer)
	var b bytes.Buffer
	fmt.Fprintln(&b, "Hello Medium") // Don't forget &
	fmt.Fprintln(os.Stdout, "Hello Godwin")

	// Optional: check the contents stored
	fmt.Println(b.String()) // Prints `HEllo Medium`

	// Writing to multiple writers at once

	// Two empty buffers
	var foo, bar bytes.Buffer

	// Create a multi writer
	mw := io.MultiWriter(&foo, &bar)

	// Write message into multi writer
	fmt.Fprintln(mw, "Hellow Ouma")

	// OPtional : verify data stored in buffers
	fmt.Println(foo.String())
	fmt.Println(bar.String())

	// Create A Simple Reader
	r := strings.NewReader("Hello Medium")

	// Read all content from reader
	c, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}

	// Optional: verify data
	fmt.Println(string(c))

}
