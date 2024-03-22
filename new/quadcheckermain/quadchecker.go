package main

import (
	"fmt"

	"os"

	"sort"

	"strings"
)

type quadFunc func(string, string) bool

type quad struct {
	name string

	f quadFunc
}

var quads = []quad{

	{"quadA", func(x, y string) bool { return x == y }},

	{"quadB", func(x, y string) bool { return x == "-"+y }},

	{"quadC", func(x, y string) bool { return x[0] == '+' && y[0] == '+' && x[1:] > y[1:] }},

	{"quadD", func(x, y string) bool { return x[0] == '-' && y[0] == '+' && x[1:] < y[1:] }},

	{"quadE", func(x, y string) bool { return x[0] == '+' && y[0] == '-' && x[1:] > "-"+y[1:] }},
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . x y")
		return
	}
	x := os.Args[1]
	y := os.Args[2]
	matches := make([]string, 0, len(quads))
	for _, q := range quads {
		if q.f(x, y) {
			matches = append(matches, fmt.Sprintf("[%s] [%s] [%s]", q.name, x, y))
		}
	}
	if len(matches) == 0 {
		fmt.Println("Not a quad function")
	} else {
		sort.Strings(matches)
		fmt.Println(strings.Join(matches, " || "))
	}
}
