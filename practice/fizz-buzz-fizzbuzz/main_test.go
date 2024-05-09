package main

import 
	"testing"
	


func TestFizzBuzz(t *testing.T) {
	 tt := []struct {
		name string
		input int
		want string
	} {
		{"Test-1", 100,"Fizz Buzz Fizz Fizz Buzz Fizz FizzBuzz Fizz Buzz Fizz Fizz Buzz Fizz FizzBuzz Fizz Buzz Fizz Fizz Buzz Fizz FizzBuzz Fizz Buzz Fizz Fizz Buzz Fizz FizzBuzz Fizz Buzz Fizz Fizz Buzz Fizz FizzBuzz Fizz Buzz Fizz Fizz Buzz Fizz FizzBuzz Fizz Buzz Fizz Fizz Buzz "},
		{"Test-2", 10, "Fizz Buzz Fizz Fizz Buzz "},
		{"Test-3", 20, "Fizz Buzz Fizz Fizz Buzz Fizz FizzBuzz Fizz Buzz "},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := GenarateFizzBuzz(tc.input)
			if got != tc.want{
				t.Errorf("%s Fail :- GenarateFizzBUzz(%d) = %s, want %s", tc.name, tc.input, got, tc.want)
			}
		})
	}
}