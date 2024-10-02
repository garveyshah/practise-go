package funcs

import (
	"errors"
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
		{"t0p", 0, errors.New("invalid charater \"t\"")},
	}

	for _, tc := range tt {
		num, err := Atoi(tc.input)

		if err != tc.err {
			t.Fatalf("for input %q - Failed :- error mismatch - got= %q, want= %q", tc.input, err, tc.err)
		}

		if num != tc.want {
			t.Fatalf("For input %q - Failed :- interger mismatch - got= %d, want= %d", tc.input, num, tc.want)
		}

	}
}

func TestIsLeapYear(t *testing.T) {
	tt := []struct {
		input int
		want  bool
	}{
		{1990, true},
		{1997, false},
		{2024, false},
	}

	for _, tc := range tt {
		got := IsLeapYear(tc.input)

		if got != tc.want {
			t.Fatalf("for input %d - Failed : bool mismatch - got= %v, want= %v", tc.input, got, tc.want)
		}
	}
}
