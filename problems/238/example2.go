package main

// @leet start
// 获取当前数的前后缀乘积
func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	result[0] = 1

	for i := 1; i < n; i++ {
		result[i] = result[i-1] * nums[i-1]
	}
	right := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= right
		right *= nums[i]
	}
	return result
}

// @leet end

func main() {}
