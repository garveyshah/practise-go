package checker

import (
	"testing"
)

func TestPrefix(t *testing.T) {
	tt := []struct {
		input []string
		want  string
	}{
		{[]string{"interview", "intermediate", "internal", "internet"}, "inter"},
		{[]string{"apple", "apricot", "app", "append"}, "ap"},
		{[]string{"station", "stable", "stand", "stair"}, "sta"},
		{[]string{"flight", "fling", "flip", "float"}, "fl"},
		{[]string{"teacher", "team", "teapot", "teal"}, "tea"},
		{[]string{"transport", "translate", "transmit", "transform"}, "trans"},
		{[]string{"sunshine", "sunday", "sunflower", "sunburn"}, "sun"},
	}

	for _, tc := range tt {
		got := Prefix(tc.input)

		if got != tc.want {
			t.Errorf("For input %q - Failed -  got= %q, want= %q", tc.input, got, tc.want)
		}
	}
}

func TestSort(t *testing.T) {
	tt := []struct {
		input []string
		want  string
	}{
		{[]string{"interview", "intermediate", "internal", "internet"}, "internal"},
		{[]string{"apple", "apricot", "app", "append"}, "app"},
		{[]string{"station", "stable", "standup", "stair"}, "stair"},
		{[]string{"flight", "fling", "flip", "float"}, "flip"},
		{[]string{"teacher", "teams", "teapot", "teal"}, "teal"},
		{[]string{"transport", "translate", "transmit", "transform"}, "transmit"},
		{[]string{"sunshine", "sunday", "sunflower", "sunburn"}, "sunday"},
	}

	for _, tc := range tt {
		got := sort(tc.input)
		if got != tc.want {
			t.Errorf("For input %q - Failed -  got= %q, want= %q", tc.input, got, tc.want)
		}
	}
}
