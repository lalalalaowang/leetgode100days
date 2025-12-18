package day8_length_of_longest_substring

func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	left, maxLength := 0, 1
	store := make(map[byte]int)

	for right := 0; right < len(s); right++ {
		if idx, ok := store[s[right]]; ok && idx >= left {
			left = idx + 1
		}
		store[s[right]] = right
		maxLength = max(maxLength, right-left+1)

	}

	return maxLength
}
