package conv

import "testing"

func TestAtoi(t *testing.T) {
	tt := []struct {
		input string
		want  int
	}{
		{"5", 5},
		{"10", 10},
		{"900", 900},
		{"-1000", -1000},
	}

	for _, tc := range tt {
		got, _ := Atoi(tc.input)
		if got != tc.want {
			t.Errorf("for input %q Failed - got- %d want= %d", tc.input, got, tc.want)
		}
	}
}
