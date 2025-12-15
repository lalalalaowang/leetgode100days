package day6_three_sum

import "sort"

func ThreeSum(nums []int) [][]int {

	sort.Ints(nums)
	result := make([][]int, 0)

	for i, num := range nums {
		if i > 0 && num == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		want := -num
		for left < right {
			if nums[left]+nums[right] == want {
				result = append(result, []int{num, nums[left], nums[right]})
				break
			} else if nums[left]+nums[right] < want {
				left++
			} else {
				right--
			}
		}
	}

	return result
}
