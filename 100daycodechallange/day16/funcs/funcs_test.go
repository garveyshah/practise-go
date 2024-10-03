package funcs

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tt := []struct {
		input1 []int
		input2 []int
		want   []int
	}{
		{[]int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
	}

	for i, tc := range tt {
		got := Merge(tc.input1, tc.input2)

		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("test [%d] failed - got= %v, want= %v", i+1, got, tc.want)
		}
	}
}

func TestSort(t *testing.T) {
	tt := []struct {
		input []int
		want  []int
	}{
		{[]int{1, 4, 5, 6, 2, 3}, []int{1, 2, 3, 4, 5, 6}},
	}

	for i, tc := range tt {
		got := Sort(tc.input)

		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("test [%d] failed - got= %v, want= %v", i+1, got, tc.want)
		}
	}
}
