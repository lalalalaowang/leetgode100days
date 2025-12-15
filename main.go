package main

import (
	"fmt"
	"leetgode100day/practices/day6_three_sum"
)

func main() {
	input := []int{-1, 0, 1, 2, -1, -4}
	result := day6_three_sum.ThreeSum(input)
	fmt.Println(result)
}
