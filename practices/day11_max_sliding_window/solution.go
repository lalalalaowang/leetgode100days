package day11_max_sliding_window

func MaxSlidingWindow(nums []int, k int) []int {
	var res []int
	dp := DeQueue{}

	for i := 0; i < k; i++ {
		dp.Push(nums[i])
	}

	res = append(res, dp.GetMax())

	for i := k; i < len(nums); i++ {
		if nums[i-k] == dp[0] {
			dp.Pop()
		}
		dp.Push(nums[i])
		res = append(res, dp.GetMax())
	}

	return res
}

type DeQueue []int

func (q *DeQueue) Push(v int) {
	for len(*q) > 0 && (*q)[len(*q)-1] < v {
		*q = (*q)[:len(*q)-1]
	}

	*q = append(*q, v)
}

func (q *DeQueue) Pop() {
	*q = (*q)[1:len(*q)]
}

func (q *DeQueue) GetMax() int {
	return (*q)[0]
}
