package day2_group_anagrams

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  [][]string
	}{
		{
			name:  "case 1",
			input: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want:  [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			name:  "case 2",
			input: []string{},
			want:  [][]string{},
		},
		{
			name:  "case 3",
			input: []string{"eat"},
			want:  [][]string{{"eat"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GroupAnagrams(tt.input)
			// 标准化实际和期望结果（对内层切片排序）
			normalizedActual := normalizeSlice(got)
			normalizedExpected := normalizeSlice(tt.want)

			// 使用 assert.ElementsMatch 忽略外层顺序
			assert.ElementsMatch(t, normalizedActual, normalizedExpected)
		})
	}
}

// normalizeSlice 对内层切片进行排序，便于比较
func normalizeSlice(slice [][]string) [][]string {
	result := make([][]string, len(slice))
	for i, inner := range slice {
		sortedInner := make([]string, len(inner))
		copy(sortedInner, inner)
		sort.Strings(sortedInner)
		result[i] = sortedInner
	}
	return result
}
