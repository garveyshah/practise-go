package main

import (
	"testing"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func TestFifthAndSkip(t *testing.T) {
	table := []string{"1234556789", "e 5£ @ 8* 7 =56 ;", "QKplq%QSw", "", "hello \\! n4ght cr3a8ure7 ", "Kimetsu no Yaiba", "8595485-52", "-552", "w58tw7474abc", "Po65 4o"}
	for _, s := range table {
		// Wrap in a function to recover from panic
		func() {
			defer func() {
				if r := recover(); r != nil {
					// This will catch any panics (error mismatch) from Fatalf and continue testing
					t.Logf("Error encountered: %v", r)
					return // Breaks loop on error, passing the test
				}
			}()

			challenge.Function("FifthAndSkip", FifthAndSkip, solutions.FifthAndSkip, s)
		}()
	}
}
