package day4_move_zeroes

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	cases := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "case1",
			input: []int{0, 1, 0, 3, 12},
			want:  []int{1, 3, 12, 0, 0},
		},
		{
			name:  "case2",
			input: []int{0},
			want:  []int{0},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if MoveZeroes(tt.input); !reflect.DeepEqual(tt.input, tt.want) {
				t.Errorf("get: %v, want: %v", tt.input, tt.want)
			}
		})
	}
}
