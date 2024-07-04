package main

import (
	"fmt"
	"testing"
	//"reflex"
)

func CustomAtoiTest(t *testing.T) {
	tt := []struct {
		name   string
		input  string
		output int
		err    error
	}{
		{"Test 1", "80", 0, nil},
		{"Test 2", "124", 123, nil},
		{"Test 3", "-123", -123, nil},
		{"Test 4", "12a3", 0, fmt.Errorf("invalid input \"a\"")},
		{"Test 5", "", 0, fmt.Errorf("invalid input: empty string")},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got, err := CustomAtoi(tc.input)
			if got != tc.output || (err != nil && err.Error() != tc.err.Error()) {
				t.Errorf("got (%d, %v), want (%d, %v)", got, err, tc.output, tc.err)

			}
		})
	}
}

func IsPrimeTest(t *testing.T) {
	tt := []struct {
		name   string
		input  int
		output string
		slice  []int
	}{
		{"Test 1", 1, "Is NOT A PRIME NUMBER", []int{1}},
		{"Test 2", 90, "Is NOT A PRIME NUMBER", []int{1, 2, 3, 5, 6, 9, 10, 15, 18, 30, 45, 90}},
		{"Test 3", -9, "Is NOT A PRIME NUMBER: It is less than 0", []int{-90}},
		{"Test 4", 97, "Is a PRIME NUMBER", []int{1, 97}},
		{"Test 5", 199, "Is a PRIME NUMBER.", []int{1, 199}},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got, slice := IsPrime(tc.input)
			if got != tc.output || !equal(slice, tc.slice) {
				t.Errorf("got (%v, %v), want (%v, %v)", got, slice, tc.output, tc.slice)
			}
		})
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
