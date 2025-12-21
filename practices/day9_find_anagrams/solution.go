package day9_find_anagrams

func FindAnagrams(s string, p string) []int {
	result := make([]int, 0)
	lenS, lenP := len(s), len(p)

	if lenP > lenS {
		return result
	}

	var storeS, storeP [26]int

	for i := 0; i < len(p); i++ {
		storeP[p[i]-'a']++
	}

	for i := 0; i < len(s); i++ {
		storeS[s[i]-'a']++
		if i >= lenP-1 {
			if storeS == storeP {
				result = append(result, i-lenP+1)
			}
			storeS[s[i-lenP+1]-'a']--
		}
	}

	return result
}
