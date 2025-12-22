package day10_subarray_sum

import (
	"reflect"
	"testing"
)

func TestSubarraySum(t *testing.T) {
	testCases := []struct {
		name   string
		input1 []int
		input2 int
		want   int
	}{
		{
			name:   "case1",
			input1: []int{1, 1, 1},
			input2: 2,
			want:   2,
		},
		{
			name:   "case2",
			input1: []int{1, 2, 3},
			input2: 3,
			want:   2,
		},
		{
			name:   "case3",
			input1: []int{-1, -1, 1},
			input2: 0,
			want:   1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := SubarraySum(tc.input1, tc.input2); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("SubarraySum() = %v, want %v", got, tc.want)
			}
		})
	}
}
