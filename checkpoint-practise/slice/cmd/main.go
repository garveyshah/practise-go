package main

import (
	"fmt"
	soln1 "project01/checkpoint-practise/slice"
)

func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", soln1.Slice(a, 1))
	fmt.Printf("%#v\n", soln1.Slice(a, 2, 4))
	fmt.Printf("%#v\n", soln1.Slice(a, -3))
	fmt.Printf("%#v\n", soln1.Slice(a, -2, -1))
	fmt.Printf("%#v\n", soln1.Slice(a, 2, 0))
}
