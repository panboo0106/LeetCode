package main

// @leet start
func isMonotonic(nums []int) bool {
	if len(nums) <= 2 {
		return true
	}
	increasing, decreasing := 0, 0
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > 0 {
			increasing = 1
		} else if nums[i]-nums[i-1] < 0 {
			decreasing = 1
		}
		if increasing == 1 && decreasing == 1 {
			return false
		}
	}
	return true
}

// @leet end

func main() {}
