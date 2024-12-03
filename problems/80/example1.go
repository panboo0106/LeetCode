package main

// @leet start
func removeDuplicates(nums []int) int {
	slow := 2
	if len(nums) < 3 {
		return len(nums)
	}
	// 只允许重复一次
	for fast := 2; fast < len(nums); fast++ {
		if nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

// @leet end

func main() {}
