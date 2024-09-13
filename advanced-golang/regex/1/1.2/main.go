// https://www.codingexplorations.com/blog/mastering-regex-pattern-matching-in-go-a-practical-guide

// An Advanced Scenario using regex for email addresses.
package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Email regex pattern with capture groups for username and domain
	emailPattern := `(?P<username>[a-zA-Z0-9._%+-]+)@(?P<domain>[a-zA-Z0-9.-]+\.[a-zA-Z]{2,})`
	emailRegex := regexp.MustCompile(emailPattern)

	// Test email
	email := "example.User+go@domain.co.uk"

	// Validate the email address
	if emailRegex.MatchString(email) {
		fmt.Println("The email address is valid.")
	} else {
		fmt.Println("The email address is not valid.")
	}

	// Extract username and domain using named capture groups
	match := emailRegex.FindStringSubmatch(email)
	result := make(map[string]string)
	for i, name := range emailRegex.SubexpNames() {
		if i > 0 && i <= len(match) {
			result[name] = match[i]
		}
	}
	fmt.Printf("Username: %s, Domain: %s\n", result["username"], result["domain"])

	// Search with a lokahead assertion: find a domain that is not followed by ".com"
	lookaheadPattern := `(?P<domain>[a-zA-Z0-9.-]+\.(?!com\b)[a-zA-Z]{2,})`
	lookaheadRegex := regexp.MustCompile(lookaheadPattern)
	domains := []string{"example.com", "domain.co.uk", "website.org"}

	for _, dom := range domains {
		if lookaheadRegex.MatchString(dom) {
			fmt.Println("Found a domain not followed by .com: %s\n")
		}
	}
}
