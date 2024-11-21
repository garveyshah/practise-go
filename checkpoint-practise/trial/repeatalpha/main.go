package main

import (
	"fmt"
	"os"
	"unicode"

)

func main() {
	testCase := []struct {
		in string
		want string
	}{
		{"abc", "abbccc"},
		{"Choumi.", "CCChhhhhhhhooooooooooooooouuuuuuuuuuuuuuuuuuuuummmmmmmmmmmmmiiiiiiiii."},
		{"", ""},
		{"abacadaba 01!", "abbacccaddddabba 01!"},
	}
	for _, tc := range testCase {
		got := RepeatAlpha(tc.in)
		if got != tc.want {
			fmt.Printf("RepeatAlpha[%q] = %q instead of %q\n", tc.in, got, tc.want)
			os.Exit(1)
		}
	}
} 

func RepeatAlpha1(s string) string {
	var ( res string )

	for _, char := range s {
		if unicode.IsLetter(char) {
			rep := unicode.ToLower(char) - 'a' + 1

			for i := 0; i < int(rep); i++ {
				res += string(char)
			}
		}   else {
			res += string(char)
		}
	}
return res
}

func RepeatAlpha(s string) string {
	var (res string; n int)

	for _, char := range s {
		if char >= 'a' &&  char <= 'z' {
			n = int(char - 'a' + 1)
		} else if char >= 'A' && char <= 'Z' {
			n = int(char - 'A' + 1)
		} else  {
			n  = 1 
		}

		for i := 1; i <= n; i++ {
			res += string(char)
		} 
	}
	return res
}