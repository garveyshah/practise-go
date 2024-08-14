package main 

import "testing"

func TestSumOfNegInts(t *testing.T) {
	 tt := []struct {
		name string
		number int
		want int
	} {
		{"Test-1", -5, -15},
		{"Test-2", -8, -36},
		{"Test-3", -5, -15},
		{"Test-4", -1000, -500500},
		{"Test-5", 0, 0}, // Test with zero input
		{"Test-6", -999999, -499999500000}, // Test with large negative number
		{"Test-7", -7, -28}, // Test with odd count of negative integers
		{"Test-8", -6, -21}, // Test with even count of negative integers
		{"Test-9", -1, -1}, // Test with negative number having absolute value of 1
		//{"Test-10", -2147483647, -2305843009213693953}, // Test with boundary case
		{"Test-11", -1000000, -500000500000},
		{"Test-12", -9223372036854775808, 0}, // Test with 64-bit boundary case (minimum value)
	//	{"Test-13", -2147483648, -4611686014132420608}, // Test with 32-bit boundary case (minimum value)
	
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := SumOfNegativeInts(tc.number)
			if  got != tc.want {
				t.Errorf("%s Failed :- SumOfNegativeInts(%d) = %d, want %d", tc.name, tc.number, got, tc.want)
			}
		})
	}
}