package main

// @leet start
func countGoodTriplets(arr []int, a int, b int, c int) int {
	// 优化思路：减少一层循环
	// 先确定第一个条件中i的大致范围
	// 在通过第二和第三个条件缩小范围
	result := 0
	prefixCount := make([]int, 1001)
	for j := 1; j < len(arr)-1; j++ {
		for i := 0; i < j; i++ {
			// 满足第一个条件的i的范围
			if abs(arr[i]-arr[j]) <= a {
				prefixCount[arr[i]]++
			}
		}
		for k := j + 1; k < len(arr); k++ {
			// 通过两个不等式确定arr[i]的取值范围
			if abs(arr[j]-arr[k]) <= b {
				left := max(arr[j]-a, arr[k]-c)
				right := min(arr[j]+a, arr[k]+c)
				result += countRange(prefixCount, left, right)
			}
		}
		// 清除技术方便下次循环计数
		for i := 0; i < j; i++ {
			if abs(arr[i]-arr[j]) <= a {
				prefixCount[arr[i]]--
			}
		}
	}
	return result
}

func countRange(prefixCount []int, left, right int) int {
	sum := 0
	left = max(0, left)
	right = min(right, 1000)
	for i := left; i <= right; i++ {
		sum += prefixCount[i]
	}
	return sum
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// @leet end

func main() {}
