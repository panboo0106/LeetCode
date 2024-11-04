package main

func moveZeroes(nums []int) {
	// 慢指针指向当前应该填入非零元素的位置
	slow := 0

	// 快指针遍历整个数组
	for fast := 0; fast < len(nums); fast++ {
		// 当遇到非零元素时
		if nums[fast] != 0 {
			// 如果两个指针不同，交换元素
			if slow != fast {
				nums[slow], nums[fast] = nums[fast], nums[slow]
			}
			// 慢指针向前移动
			slow++
		}
	}
}

func main() {}
