package sort

import "math/rand"

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
		if !flag {
			break
		}
	}
	return arr
}

// 插入排序实现
// 默认数组一个元素已经排好序，剩余其他元素依次作为pivot元素根据大小排序
// 插入到已排序区间内

func InsertSort(arr []int) []int {
	// 已排序区间设为 [0, i-1]
	for i := 1; i < len(arr); i++ {
		// 设置base元素
		pivot := arr[i]
		// 依次从右到左比较已排序区间元素
		for j := i; j > 0 && arr[j-1] > pivot; j-- {
			arr[j] = arr[j-1]
		}
	}
	return arr
}

// 快速排序实现
// 双边循环法 Hoare Partition
// 逻辑：把数组划分为小于等于大于pivot区间
// 默认选择第一个元素为pivot元素，并且在数组开头和末尾设置两个指针
// 分别向左（向右）寻找大于（小于）pivot元素，找到后交换这两个元素
// 并移动左右指针到对应元素，直到左右指针相遇。把该位置元素与pivot元素进行交换到此完成第一轮迭代
// 后续把pivot元素左右划分成两部分分别设置pivot元素再做上面相同的操作

func QuickSort(arr []int, left, right int) []int {
	if left >= right {
		return arr
	}
	pivot := partition(arr, left, right)
	// 通过递归的方式，这里传入切片
	// 为防止尾递归问题，判断区间范围，范围小的先执行
	if pivot-left > right-pivot {
		QuickSort(arr, pivot+1, right)
	} else {
		QuickSort(arr, left, pivot-1)
	}
	return arr
}

func partition(arr []int, left, right int) int {
	// 为防止基数影响随机取值与第一个元素进行交换，把arr[l]设置为pivot元素
	pivot := left + rand.Intn(right-left+1)
	arr[left], arr[pivot] = arr[pivot], arr[left]
	i, j := left, right
	for i < j {
		// 从左往右遍历找到大于pivot元素
		for i < j && arr[i] < arr[left] {
			i++
		}
		// 从右到左遍历找到小于pivot元素
		for i < j && arr[j] >= arr[left] {
			j--
		}
		// 交换左右指针的元素
		arr[i], arr[j] = arr[j], arr[i]

	}
	// 把pivot元素与当前位置元素进行交换
	arr[i], arr[left] = arr[left], arr[i]
	// 返回划分点
	return i
}

// 单边循环法 lomuto Partition
// 逻辑：把数组划分为小于pivot和大于等于pivot的两个区间
// (基准元素可以放在数组第一位或者最后一位)
// 1.基准元素放在第一位时
// 设置一个基准元素和一个指针指向数组起始位置,从基准元素的下一个元素开始判断，
// 2.基准元素放在最后一位时
// 从左到右开始判断
// 如果小于基准元素指针向前移动一位,代表小于基准元素的区间又扩大一位
// 随后交换循环所在元素与指针所在元素，单次循环结束后交换pivot和指针所在元素

func lomutoPartition_first(arr []int, left, right int) int {
	// 设置随机pivot元素，并把pivot元素与数组第一个元素交换
	pivotIndex := left + rand.Intn(right-left+1)
	arr[left], arr[pivotIndex] = arr[pivotIndex], arr[left]

	// 设置mark标记记录小于pivot元素区间
	mark := left

	for i := left + 1; i <= right; i++ {
		if arr[i] < arr[left] {
			mark++
			arr[mark], arr[i] = arr[i], arr[mark]
		}
	}
	// 遍历之后mark指针左侧区间元素都小于pivot元素
	// 交换mark指针所在元素和pivotIndex所在元素
	arr[left], arr[mark] = arr[mark], arr[left]
	return mark
}

// 非递归方式实现
// 定义栈存储数组范围
type Stack struct {
	items [][2]int
}

func (s *Stack) Push(item [2]int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() ([2]int, bool) {
	if len(s.items) == 0 {
		return [2]int{}, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func QuickSortNonR(arr []int) {
	if len(arr) <= 1 {
		return
	}

	s := &Stack{}
	s.Push([2]int{0, len(arr) - 1})
	for !s.IsEmpty() {
		range_, ok := s.Pop()
		if !ok {
			continue
		}
		low, high := range_[0], range_[1]
		if low < high {
			pivotIndex := lomutoPartition_first(arr, low, high)

			// 将较大的子数组范围先压入栈，以减少栈的使用(范围越大意味着入栈次数越少)
			if pivotIndex-low < high-pivotIndex {
				s.Push([2]int{pivotIndex + 1, high})
				s.Push([2]int{low, pivotIndex - 1})
			} else {
				s.Push([2]int{low, pivotIndex - 1})
				s.Push([2]int{pivotIndex + 1, high})
			}
		}
	}
}
