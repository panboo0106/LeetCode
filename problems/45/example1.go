package main

// 不要纠结于具体跳到哪个位置，而是关注在当前能够到达的范围内，
// j哪里能作为"跳板"让我们下一步跳得更远
// @leet start
func jump(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	maxReach := 0
	jumpCount := 0
	currEnd := 0
	for i := 0; i < len(nums)-1; i++ {
		// 更新当前位置能达到的最远距离
		maxReach = max(maxReach, nums[i]+i)
		// 如果达到当前可跳跃的最大范围，必须进行下一次跳跃
		// 请注意这里的跳跃是在比较所有的范围的可到达
		// 下一跳的结果后进行的，不是单纯就是当前位置
		// 最大可跳跃长度
		if i == currEnd {
			jumpCount++
			currEnd = maxReach
		}
	}
	return jumpCount
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @leet end

func main() {}
