func canJump(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	maxReach := nums[0] // 最远可到达距离
	for i := 0; i < len(nums); i++ {
		if i > maxReach {
			return false
		}
		maxReach = max(maxReach, i+nums[i])
		if maxReach >= len(nums)-1 {
			return true
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

