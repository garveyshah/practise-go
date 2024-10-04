package compute

import "testing"

func TestCompute(t *testing.T) {
	tt := []struct {
		input int
		want  int
	}{
		{1, 1},
		{5, 5},
		{11, 89},
	}

	for _, tc := range tt {
		got := Fibonnaci(tc.input)
		if tc.want != got {
			t.Errorf("For input %d Failed - got= %d, want= %d", tc.input, got, tc.want)
		}
	}
}
