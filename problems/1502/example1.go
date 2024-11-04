package main

import "math/rand"

// @leet start
func canMakeArithmeticProgression(arr []int) bool {
	// 先排序
	if len(arr) <= 2 {
		return true
	}
	left, right := 0, len(arr)-1
	QuickSort(arr, left, right)
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != arr[1]-arr[0] {
			return false
		}
	}
	return true
}

func QuickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	pivot := partition(arr, left, right)
	if pivot-left >= right-pivot {
		QuickSort(arr, pivot+1, right)
		QuickSort(arr, left, pivot-1)
	} else {
		QuickSort(arr, left, pivot-1)
		QuickSort(arr, pivot+1, right)
	}
}

func partition(arr []int, left, right int) int {
	// 为防止基数影响随机取值与第一个元素进行交换，把arr[l]设置为pivot元素
	pivot := left + rand.Intn(right-left+1)
	arr[left], arr[pivot] = arr[pivot], arr[left]
	i, j := left, right
	for i < j {
		// 从左往右遍历找到大于等于pivot元素
		for i < j && arr[i] < arr[left] {
			i++
		}
		// 从右到左遍历找到小于pivot元素
		for i < j && arr[j] >= arr[left] {
			j--
		}
		if i < j {
			// 交换左右指针的元素
			arr[i], arr[j] = arr[j], arr[i]
		}

	}
	// 把pivot元素与当前位置元素进行交换
	arr[i], arr[left] = arr[left], arr[i]
	// 返回划分点
	return i
}

// @leet end

func main() {}
