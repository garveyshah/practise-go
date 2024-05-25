package sorting

import (
	"reflect"
	"testing"
)

func TestBubbleSort1(t *testing.T) {
	tt := []struct {
		name  string
		input []int
		want  []int
	}{
		{"test 1", []int{3, 2, 1}, []int{1, 2, 3}},
		{"test 2", []int{5, 1, 4, 2, 8}, []int{1, 2, 4, 5, 8}},
		{"Test 3", []int{5, 100, 7, 13, 78, 50, 6, 23, 90}, []int{5, 6, 7, 13, 23, 50, 78, 90, 100}},
		{"Test 4", []int{1000, 999, 567, 6, 65, 89, 876, 546, 287}, []int{6, 65, 89, 287, 546, 567, 876, 999, 1000}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			BubbleSort(tc.input)
			if !reflect.DeepEqual(tc.input, tc.want) {
				t.Errorf("%v Failed :- got %v, want %v", tc.name, tc.input, tc.want)
			}
		})
	}
}
