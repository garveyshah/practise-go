package main

import (
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {
	// Test cases
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"Test 1", []int{4, 5, 6, 4, 5, 7}, []int{4, 5, 6, 7}},
		{"Test 2", []int{1}, []int{1}},
		{"Test 3", []int{}, []int{}},
		{"Test 4", []int{2, 2, 2, 2, 2}, []int{2}},
		{"Test 5", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"Test 6", []int{1, 2, 3, 2, 1, 3, 4}, []int{1, 2, 3, 4}},
		{"Test 7", []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}, []int{1, 2, 3, 4, 5}},
		{"Test 8", []int{-1, -2, -3, -1, -2, -4}, []int{-1, -2, -3, -4}},
		{"Test 9", []int{1, -1, 2, -2, 1, -1}, []int{1, -1, 2, -2}},
		{"Test 10", []int{0, 1, 0, 2, 0, 3}, []int{0, 1, 2, 3}},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := RemoveDups(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("FAILED: input = %v, expected = %v, got = %v\n", tc.input, tc.expected, result)
			}
		})
	}
}
