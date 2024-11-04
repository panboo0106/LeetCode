package main

// @leet start
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	// 处理全是9的情况
	result := make([]int, len(digits)+1)
	result[0] = 1
	return result
}

// @leet end

func main() {}
