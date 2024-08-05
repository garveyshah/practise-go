package main

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tt := []struct {
		name     string
		input    []int
		expected []int
	}{
		{name: "Test 1", input: []int{7}, expected: []int{7}},                                     // Single element array
		{name: "Test 2", input: []int{}, expected: []int{}},                                       // Empty array
		{name: "Test 3", input: []int{5, 5, 5, 5, 5}, expected: []int{5}},                         // All elements are the same
		{name: "Test 4", input: []int{1, 2, 3, 4, 5}, expected: []int{1, 2, 3, 4, 5}},             // No duplicates
		{name: "Test 5", input: []int{1, 2, 2, 3, 4, 4, 5}, expected: []int{1, 2, 3, 4, 5}},       // Some duplicates
		{name: "Test 6", input: []int{-1, -1, -2, -2, -3}, expected: []int{-1, -2, -3}},           // Negative numbers
		{name: "Test 7", input: []int{-1, 1, -2, 2, -3, 3}, expected: []int{-1, 1, -2, 2, -3, 3}}, // Mixed positive and negative numbers
		{name: "Test 8", input: []int{0, 0, 0, 1, 2, 3}, expected: []int{0, 1, 2, 3}},             // Array with zero and other numbers
	}

	for _, tc := range tt {
		result := RemoveDuplicates(tc.input)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("%v - removeDuplicates(%v) = %v; expected %v",tc.name, tc.input, result, tc.expected)
		}
	}
}