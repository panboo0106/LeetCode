package main

// @leet start
func removeElement(nums []int, val int) int {
	// 慢指针
	slow := 0
	for fast := range nums {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

// @leet end

func main() {}
