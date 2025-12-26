package day12_min_window

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	left := 0
	resLeft := 0
	resRight := 0
	storeT := [128]int{}
	storeS := [128]int{}

	for i := 0; i < len(t); i++ {
		storeT[t[i]]++
	}

	for i := len(t) - 1; i >= 0; i-- {
		storeS[t[i]]++
		if storeS == storeT {
			if (resLeft == 0 && resRight == 0) || (i-left < resRight-resLeft) {
				resLeft = left
				resRight = i
			}

			storeS[left]--
			left++
		}
	}

	return s[resLeft : resRight+1]
}
