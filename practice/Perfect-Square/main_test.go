package main

import (
	"testing"
)

func TestPerfectSquare(t *testing.T) {
	 tt := [] struct {
		name string
		input int
		want string
	}{
		{"Test-1", 25,"a Perfect Square"},
		{"Test-2", 98,"not a Perfect Square"},
		{"Test-3", 10000,"a Perfect Square"},
		{"Test-4", 100,"a Perfect Square"},
		{"Test-5", 625,"a Perfect Square"},
		{"Test-6", 47,"not a Perfect Square"},
		{"Test-7", 81,"a Perfect Square"},
		{"Test-8", 1000000,"a Perfect Square"},
		{"Test-9", 390625,"a Perfect Square"},
		{"Test-10", 907,"not a Perfect Square"},
		{"Test-11", 255,"not a Perfect Square"},
		{"Test-12", 2343750,"not a Perfect Square"},
		{"Test-13", 11718750,"not a Perfect Square"},
		{"Test-14", 36076,"not a Perfect Square"},
		{"Test-15", 1875952,"not a Perfect Square"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := 	PerfectSquare(tc.input) 
				if got != tc.want {
					t.Errorf("%s Failed, PerfectSquare(%d) = %s, want %s.", tc.name, tc.input, got, tc.want)
				}
			})
		}
	}
