package main

// @leet start
func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n // 处理k大于数组长度的情况
	if k == 0 {
		return
	}

	// 1.反转整个数组
	reverse(nums, 0, n-1)
	// 2.反转前k个
	reverse(nums, 0, k-1)
	// 3.反转剩余的
	reverse(nums, k, n-1)
}

// 反转函数
func reverse(nums []int, start, end int) {
	for start < end {
		// 交换start和end位置的元素
		nums[start], nums[end] = nums[end], nums[start]
		// 移动指针
		start++
		end--
	}
}

// @leet end

func main() {}
