package golang

// 选择排序实现
// 开启一个循环，每轮从未排序区间选择最小的元素，将其放到已排序区间的末尾
func SelectSort(arr []int) []int {
	// 获取未排序区间内的第一个元素
	for i := 0; i < len(arr)-1; i++ {
		k := i
		// 找到未排序区间内最小的一个元素，把该元素放到已排序区间最后
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
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
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j+1] < arr[j] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
	return arr
}
