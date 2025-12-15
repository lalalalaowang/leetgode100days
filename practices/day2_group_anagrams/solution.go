package day2_group_anagrams

import (
	"sort"
)

func GroupAnagrams(input []string) [][]string {
	store := make(map[string][]string)

	for _, str := range input {
		runes := []rune(str)

		sort.Slice(runes, func(i, j int) bool {
			return runes[i] > runes[j]
		})

		title := string(runes)

		if _, ok := store[title]; !ok {
			store[title] = make([]string, 0)
		}

		store[title] = append(store[title], str)
	}

	result := make([][]string, 0)

	for _, v := range store {
		result = append(result, v)
	}

	return result
}
