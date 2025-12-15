package day5_max_area

func MaxArea(height []int) int {
	left := 0
	right := len(height) - 1
	result := 0

	for left != right {
		result = maxInt(minInt(height[left], height[right])*(right-left), result)

		if height[left] > height[right] {
			right -= 1
		} else {
			left += 1
		}
	}

	return result
}

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func minInt(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
