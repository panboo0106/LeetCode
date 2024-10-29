package main

// @leet start
func strStr(haystack string, needle string) int {
	needleLen := len(needle)
	start, end := 0, 0+needleLen
	for end < len(haystack) {
		if haystack[start:end] == needle {
			return start
		} else {
			start++
			end = start + needleLen
		}
	}
	return -1
}

// @leet end

func main() {}
