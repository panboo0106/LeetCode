package main

// @leet start
func moveZeroes(nums []int) {
	// 从后往前
	lenght := len(nums)
	// 记录非0位置指针
	j := lenght - 1
	for i := lenght - 1; i >= 0; i-- {
		if nums[i] == 0 {
			temp := i
			// 找到零后逐个交换位置
			for temp < j {
				nums[temp], nums[temp+1] = nums[temp+1], nums[temp]
				temp++
			}
			// 慢指针向前移动
			j--

		} else {
			continue
		}
	}
	return
}

// @leet end

func main() {}
