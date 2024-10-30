package main

// @leet start
func repeatedSubstringPattern(s string) bool {
	pi := make([]int, len(s))
	n := len(s)
	for i := 1; i < len(s); i++ {
		length := pi[i-1]
		for length > 0 && s[i] != s[length] {
			length = pi[length-1]
		}
		if s[i] == s[length] {
			length++
		}
		pi[i] = length
	}
	last := pi[len(s)-1]
	// 如果字符串长度能被 (n-last) 整除，且last不为0
	return last > 0 && n%(n-last) == 0
}

// @leet end

func main() {}
