package main

// @leet start
func majorityElement(nums []int) int {
	// 候选数和计数器
	candidate := nums[0]
	count := 1

	// 遍历数组
	for i := 1; i < len(nums); i++ {
		// 计数器为0时,更换候选数
		if count == 0 {
			candidate = nums[i]
		}

		// 当前数等于候选数,计数加1
		// 否则计数减1
		if nums[i] == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}

// @leet end

func main() {}
