package day10_subarray_sum

func SubarraySum(nums []int, k int) int {
	result, pre := 0, 0

	if len(nums) == 0 {
		return result
	}

	store := make(map[int]int)
	store[0] = 1

	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := store[pre-k]; ok {
			result += store[pre-k]
		}
		store[pre] += 1
	}

	return result
}
