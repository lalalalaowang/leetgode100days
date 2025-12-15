package day3_longest_consecutive

func LongestConsecutive(nums []int) int {
	numMap := make(map[int]bool)

	for _, v := range nums {
		numMap[v] = true
	}

	longest := 0

	for k, _ := range numMap {
		long := 0
		if !numMap[k-1] {
			currentNum := k
			for numMap[currentNum] {
				long++
				currentNum++
			}
		}
		if long > longest {
			longest = long
		}
	}

	return longest
}
