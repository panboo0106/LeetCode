package main

// @leet start
func removeDuplicates(nums []int) int {
	slow := 0
	for fast := range nums {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}

// @leet end

func main() {}
