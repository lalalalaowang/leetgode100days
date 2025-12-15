package day1_two_sum

func TwoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, v := range nums {
		other := target - v
		if idx, ok := numMap[other]; ok {
			return []int{idx, i}
		}
		numMap[v] = i
	}

	return nil
}
