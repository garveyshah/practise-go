package main

import "testing"

func TestSumOfPosInts(t *testing.T) {
	tt := []struct {
		name   string
		number int
		want   int
	}{
		{"Test-1", 5, 15},
		{"Test-2", 0, 0},
		{"Test-3", 1000000, 500000500000},
		{"Test-4", 10000, 50005000},
		{"Test-5", -5, 0},
		{"Test-6", 55, 1540},
		{"Test-7", 2147483647, 2305843008139952128},
		{"Test-8", 1, 1},
		{"Test-9", 2147483647, 9223372036854775807},
		{"Test-10", 1, 1},
		{"Test-11", 9223372036854775807, 9223372036854775807},
		{"Test-12", 9223372036854775807, 427611669758975826},
		{"Test-13", 999999, 499999500000},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := SumOfPositiveInts(tc.number)
			if got != tc.want {
				t.Errorf("%s Failed! :- SumOfPositiveInts(%d) = %d, want %d", tc.name, tc.number, got, tc.want)
			}
		})
	}
}
