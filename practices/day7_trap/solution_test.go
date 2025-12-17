package day7_trap

import (
	"reflect"
	"testing"
)

func TestTrap(t *testing.T) {
	testCases := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "case1",
			input: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			want:  6,
		},
		{
			name:  "case2",
			input: []int{4, 2, 0, 3, 2, 5},
			want:  9,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if get := Trap(tc.input); !reflect.DeepEqual(tc.want, get) {
				t.Errorf("want: %v, got: %v", tc.want, get)
			}
		})
	}
}
