package main

// @leet start
func strStr(haystack string, needle string) int {
	n, m := len(needle), len(haystack)

	// 处理边界情况
	if n == 0 {
		return 0
	}
	if n > m {
		return -1
	}

	// KMP算法实现
	// 目的就是通过前缀函数判断模式串是否匹配
	s := needle + "#" + haystack
	pi := make([]int, len(s))

	for i := 1; i < len(s); i++ {
		length := pi[i-1]
		for length > 0 && s[i] != s[length] {
			length = pi[length-1]
		}
		if s[i] == s[length] {
			length++
		}
		pi[i] = length
		if pi[i] == n {
			return i - 2*n
		}
	}
	return -1
}

// @leet end

func main() {}
