package main

// @leet start
func rotate(nums []int, k int) {
	// 使用双指针构建循环数组
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		result[(i+k)%len(nums)] = nums[i]
	}
	copy(nums, result)
}

// @leet end

func main() {}
