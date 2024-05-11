package main

import "testing"

func TestLargerstPrimerFact(t *testing.T) {
	tt := []struct {
		name  string
		input int
		want  int
	}{
		{"Test-1", 183486534, 1237},
		{"Test-2", 367908073, 2593},
		{"Test-3", 549702467, 11927},
		{"Test-4", 43300748, 197},
		{"Test-5", 1407363469, 21787},
		{"Test-6", 215309001, 13099},
		{"Test-7", 237671281, 1987},
		{"Test-8", 570129211, 14281},
		{"Test-9", 353894333, 17093},
		{"Test-10", 260006285, 3433},
		{"Test-11", 209627975, 18397},
		{"Test-12", 24622481, 1873},
		{"Test-13", 125803543, 2039},
		{"Test-14", 589805651, 1531},
		{"Test-15", 589805651, 1531},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := LargestPrimeFact(tc.input)
			if got != tc.want {
				t.Errorf("%s, FAiled :- LargestPrimeFact(%d) = %d, want %d.", tc.name, tc.input, got, tc.want)
			}
		})
	}
}
