package day8_length_of_longest_substring

func LengthOfLongestSubstring(s string) int {
	result := 0
	length := 0
	runes := []rune(s)

	store := make(map[rune]int)

	for i, item := range runes {
		length = result
		if _, ok := store[item]; ok {
			if length > result {
				result = length
			}
		} else {
			store[item] = i
			length += 1
		}
	}

	if length > result {
		result = length
	}

	return result
}
