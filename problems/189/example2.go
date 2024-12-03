package main

// @leet start
func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n // 处理k超过数组长度的情况
	if k == 0 {
		return
	}

	count := 0 // 记录已经处理的元素个数
	start := 0 // 当前处理的起始位置

	for count < n {
		current := start    // 当前处理的位置
		prev := nums[start] // 当前位置的值

		for {
			// 计算下一个位置
			next := (current + k) % n
			// 保存下一个位置的值
			temp := nums[next]
			// 将当前值放入下一个位置
			nums[next] = prev
			// 更新prev为被替换的值
			prev = temp
			// 更新当前位置
			current = next
			// 处理元素计数加1
			count++

			// 如果回到起始位置，结束当前环
			if start == current {
				break
			}
		}
		// 如果还有元素未处理，开始处理下一个环
		start++
	}
}

// @leet end

func main() {}
