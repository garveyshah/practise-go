package funcs

import (
	"fmt"
	"testing"
)

func TestAtoi(t *testing.T) {
	tt := []struct {
		input string
		want  int
		err   error
	}{
		{"202", 202, nil},
		{"3009", 3009, nil},
		{"0", 0, nil},
		{"t0p", 0, fmt.Errorf("invalid character \"t\"")},
	}

	for _, tc := range tt {
		num, err := Atoi(tc.input)

		if err == nil && tc.err != nil || err != nil && tc.err == nil || (err != nil && err.Error() != tc.err.Error()) {
			t.Errorf("for input %q - Failed :- error mismatch - got= %s, want= %s", tc.input, err, tc.err)
		}

		if num != tc.want {
			t.Errorf("For input %q - Failed :- interger mismatch - got= %d, want= %d", tc.input, num, tc.want)
		}

	}
}

func TestIsLeapYear(t *testing.T) {
	tt := []struct {
		input int
		want  bool
	}{
		{2024, true},
		{2032, true},
		{2050, false},
		{300, true},
		{2000, true},
		{2022, false},
		{1994, false},
		{2016, true},
		{1900,true},
	}

	for _, tc := range tt {
		got := IsLeapYear(tc.input)

		if got != tc.want {
			t.Errorf("for input %d - Failed : bool mismatch - got= %v, want= %v", tc.input, got, tc.want)
		}
	}
}
