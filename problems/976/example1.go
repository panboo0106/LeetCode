package main

import "sort"

// @leet start
func largestPerimeter(nums []int) int {
	sort.Ints(nums)
	for i := len(nums) - 1; i > 1; i-- {
		for j := i - 1; j > 0; j-- {
			for k := j - 1; k >= 0; k-- {
				if nums[k]+nums[j] > nums[i] {
					return nums[k] + nums[j] + nums[i]
				} else {
					break
				}
			}
		}
	}
	return 0
}

// @leet end

func main() {}
