package base

import "math/rand"

func randomAccess(nums []int) int {
	// 随机抽取一个数
	randomIndex := rand.Intn(len(nums))
	// 返回随机数
	return nums[randomIndex]
}

func insert(arr []int, num, index int) []int {
	// 判断输入索引是否正确
	if index < 0 || index > len(arr) {
		return arr
	}
	// 插入数组意味着有旧数据丢失，通过先扩展在插入方式
	arr = extend(arr)
	// 找到插入位置后，之前元素向后移动一位
	for i := len(arr) - 1; i > index; i-- {
		arr[i] = arr[i-1]
	}
	// 使用浅拷贝更优
	// copy(arr[index+1:], arr[index:])
	arr[index] = num
	return arr
}

func extend(arr []int) []int {
	// 用 0 填充数组
	return append(arr, 0)
}

func delete(arr []int, index int) []int {
	// 判断输入索引是否正确
	if index < 0 || index > len(arr) {
		return arr
	}
	// 之后元素向前移动一位
	for i := index; i < len(arr)-1; i++ {
		arr[i] = arr[i+1]
	}
	return arr
}

// 查找
func find(arr []int, target int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return true
		}
	}
	return false
}
