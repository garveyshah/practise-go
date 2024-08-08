package main

import "testing"

// Test cases for HashCode function
func TestHashCode(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello", "Mjqqt"},
		{"World", "|twqi"},
		{"12345", "56789"},
		{"!@#$%", "-./01"},
		{"", ""},
		{"GoLang", "LqQfrl"},
		{"a quick brown fox", "e uymeo%fvsarjsz|"},
		{"The quick brown fox jumps over the lazy dog", "Xli%uymeq$fvsarjsz$nyrtw%sziw%xli$pdez$hs"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ", "EFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcd"},
		{"abcdefghijklmnopqrstuvwxyz", "efghijklmnopqrstuvwxyz{|}~ !\"#"},
		{"0123456789", "456789:;<=>?"},
		{"!@#$%^&*()_+-=[]{}|;':,./<>?`~", "-./01BCD,-02345WXYZ7:<<5=>[7]"},
	}

	for _, test := range tests {
		result := HashCode(test.input)
		if result != test.expected {
			t.Errorf("HashCode(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}
