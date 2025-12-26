package day12_min_window

func minWindow(s string, t string) string {
	window := make(map[byte]int, 0)
	need := make(map[byte]int, 0)

	left, match := -1, 0
	start, end, minLen := 0, 0, len(s)+1

	for i := range t {
		need[t[i]]++
	}

	for right := 0; right < len(s); right++ {
		ch1 := s[right]
		window[ch1]++

		if window[ch1] == need[ch1] {
			match++
		}

		for match == len(need) {
			if right-left < minLen {
				start, end = left, right
				minLen = right - left
			}

			left++
			ch2 := s[left]
			if window[ch2] == need[ch2] {
				match--
			}
			window[ch2]--
		}
	}

	return s[start+1 : end+1]
}
