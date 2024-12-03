package main

import (
	"slices"
)

// @leet start
func removeElement(nums []int, val int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums[i] = 1000
		} else {
			result++
		}
	}
	slices.Sort(nums)
	return result
}

// @leet end

func main() {}
