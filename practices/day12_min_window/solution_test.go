package day12_min_window

import (
	"reflect"
	"testing"
)

func TestMinWindow(t *testing.T) {
	testCases := []struct {
		name   string
		input1 string
		input2 string
		want   string
	}{
		{
			name:   "case1",
			input1: "ADOBECODEBANC",
			input2: "ABC",
			want:   "BANC",
		},
		{
			name:   "case2",
			input1: "a",
			input2: "a",
			want:   "a",
		},
		{
			name:   "case3",
			input1: "a",
			input2: "aa",
			want:   "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := minWindow(tc.input1, tc.input2); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("minWindow() = %v, want %v", got, tc.want)
			}
		})
	}
}
