package main

import (
	"fmt"
	"os"
	"unicode"
)

func RepeatAlpha(str string) string {
	// res := ""
	// count := 0

	// for _, char := range str {
	// 	if char >= 'A' && char <= 'Z' {
	// 		count = int(char - 'A' + 1)
	// 	} else if char >= 'a' && char <= 'z' {
	// 		count = int(char - 'a' + 1)
	// 	} else {
	// 		count = 1
	// 	}

	// 	for i := 0; i < count; i++ {
	// 		res += string(char)
	// 	}
	// }
	// return res

	res := ""

	for _, r := range str {
		if unicode.IsLetter(r) {
			rep := unicode.ToLower(r) - 'a' + 1
			for i := 0; i < int(rep); i++ {
				res += string(r)
			}
		} else {
			res += string(r)
		}
	}
	return res
}

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))

	testCases := []struct {
		in   string
		want string
	}{
		{"abc", "abbccc"},
		{"Choumi.", "CCChhhhhhhhooooooooooooooouuuuuuuuuuuuuuuuuuuuummmmmmmmmmmmmiiiiiiiii."},
		{"", ""},
		{"abacadaba 01!", "abbacccaddddabba 01!"},
	}
	for _, tc := range testCases {
		got := RepeatAlpha(tc.in)
		if got != tc.want {
			fmt.Printf("RepeatAlpha(%q) = %q instead of %q\n", tc.in, got, tc.want)
			os.Exit(1)
		}
	}
}
