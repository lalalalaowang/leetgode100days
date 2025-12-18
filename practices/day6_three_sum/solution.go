package day6_three_sum

import (
	"sort"
	"strconv"
)

func ThreeSum(nums []int) [][]int {

	sort.Ints(nums)
	store := make(map[string][]int)

	for i, num := range nums {
		if i > 0 && num == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		want := -num
		for left < right {
			if nums[left]+nums[right] == want {
				if _, ok := store[strconv.Itoa(num)+strconv.Itoa(nums[left])+strconv.Itoa(nums[right])]; !ok {
					store[strconv.Itoa(num)+strconv.Itoa(nums[left])+strconv.Itoa(nums[right])] = []int{num, nums[left], nums[right]}
				}
				left++
				right--
			} else if nums[left]+nums[right] < want {
				left++
			} else {
				right--
			}
		}
	}

	result := make([][]int, 0)
	for _, num := range store {
		result = append(result, num)
	}

	return result
}
