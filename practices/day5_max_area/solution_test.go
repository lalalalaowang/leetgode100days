package day5_max_area

import (
	"reflect"
	"testing"
)

func TestMaxArea(t *testing.T) {
	cases := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "case1",
			input: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:  49,
		},
		{
			name:  "case2",
			input: []int{1, 1},
			want:  1,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if get := MaxArea(tt.input); !reflect.DeepEqual(get, tt.want) {
				t.Errorf("get: %v, want: %v", get, tt.want)
			}
		})
	}
}
