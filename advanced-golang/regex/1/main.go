// https://www.codingexplorations.com/blog/mastering-regex-pattern-matching-in-go-a-practical-guide

/* Mastering Regex Pattern Matching in Go: A Practical Guide */
/* Setting up Regex in Go
- First import "regexp"
- The start by creating a regex object with regexp.Compile.
- Or regexp.MustCompile - Panics if the expression cannot be parsed,
 which is useful when you  want to ensure the pattern is valid at commpile time*/

package main

import (
	"fmt"
	"regexp"
)

func main() {
	r, err := regexp.Compile(`\d{3}-\d{2}-\d{4}`)
	if err != nil {
		fmt.Println("Errror != nil: ", err)
		return
	}
	fmt.Println(r, "\n")

	/* Performaing Pattern Matching */
	// Matching
	matched := r.MatchString("123-45-6789")
	fmt.Println(matched, "\n") // Outputs: true

	// Finding first occurencece of a pattern:
	match := r.FindString("My SSN is 123-45-6789.")
	fmt.Println(match, "\n") // Outputs: 123-45-

	// Finding All Matches - finding all occurrences of a pattern:
	matches := r.FindAllString("123-45-6789 and 987-65-4321", -1)
	fmt.Println(matches, "\n") // Outputs: [123-45-6789 987-65-4321]

	// Replacing - repalce matches with another string:
	replaced := r.ReplaceAllString("SSN: 123-45-6789", "REDACTED")
	fmt.Println(replaced, "\n") // Outputs: SSN: REDACTED
}
