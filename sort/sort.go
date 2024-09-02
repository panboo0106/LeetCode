package sort

// 选择排序实现
// 开启一个循环，每轮从未排序区间选择最小的元素，将其放到已排序区间的末尾
func SelectSort(arr []int) []int {
	// 获取未排序区间内的第一个元素
	for i := 0; i < len(arr)-1; i++ {
		k := i
		// 找到未排序区间内最小的一个元素，把该元素放到已排序区间最后
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[k] {
				k = j
			}
		}
		arr[i], arr[k] = arr[k], arr[i]
	}
	return arr
}

// 冒泡排序实现
// 遍历数组，依次比较相邻两个数，如果左元素大于右元素就交换，
// 直到每轮中把该轮循环中最大的数放到该轮的最右边

func BubbleSort(arr []int) []int {
	// 未排序区间 [0 i]
	for i := len(arr) - 1; i > 0; i-- {
		flag := false // 初始化标志位
		// 在未排序区间内依次比较相邻元素
		for j := 0; j < i; j++ {
			if arr[j+1] < arr[j] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				flag = true
			}
		}
		if flag == false {
			break
		}
	}
	return arr
}

// 插入排序实现
// 默认数组一个元素已经排好序，剩余其他元素依次作为base元素根据大小排序
// 插入到已排序区间内

func InsertSort(arr []int) []int {
	// 已排序区间设为 [0, i-1]
	for i := 1; i < len(arr); i++ {
		// 设置base元素
		base := arr[i]
		// 依次从右到左比较已排序区间元素
		for j := i; j > 0 && arr[j-1] > base; j-- {
			arr[j] = arr[j-1]
		}
	}
	return arr
}
