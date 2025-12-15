package day4_move_zeroes

func MoveZeroes(nums []int) {
	slow := 0

	for fast, num := range nums {
		if num != 0 {
			nums[fast], nums[slow] = nums[slow], nums[fast]
			slow++
		}
	}
}
