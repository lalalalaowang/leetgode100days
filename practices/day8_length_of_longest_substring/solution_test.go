package day8_length_of_longest_substring

import (
	"reflect"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	teatCases := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "case1",
			input: "abcabcbb",
			want:  3,
		},
		{
			name:  "case2",
			input: "bbbbb",
			want:  1,
		},
		{
			name:  "case3",
			input: "pwwkew",
			want:  3,
		},
	}

	for _, tc := range teatCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := LengthOfLongestSubstring(tc.input); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("LengthOfLongestSubstring() = %v, want %v", got, tc.want)
			}
		})
	}
}
