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
