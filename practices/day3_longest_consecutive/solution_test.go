package day3_longest_consecutive

import (
	"reflect"
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	cases := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "case1",
			input: []int{100, 4, 200, 1, 3, 2},
			want:  4,
		},
		{
			name:  "case2",
			input: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			want:  9,
		},
		{
			name:  "case3",
			input: []int{1, 0, 1, 2},
			want:  3,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestConsecutive(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
