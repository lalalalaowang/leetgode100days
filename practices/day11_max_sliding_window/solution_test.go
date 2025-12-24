package day11_max_sliding_window

import (
	"reflect"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	testCases := []struct {
		name   string
		input1 []int
		input2 int
		want   []int
	}{
		{
			name:   "case1",
			input1: []int{1, 3, -1, -3, 5, 3, 6, 7},
			input2: 3,
			want:   []int{3, 3, 5, 5, 6, 7},
		},
		{
			name:   "case2",
			input1: []int{1},
			input2: 1,
			want:   []int{1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := MaxSlidingWindow(tc.input1, tc.input2); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("maxSlidingWindow() = %v, want %v", got, tc.want)
			}
		})
	}
}
