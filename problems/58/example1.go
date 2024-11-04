package main

// @leet start
func lengthOfLastWord(s string) int {
	result := 0
	flag := false
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) != " " {
			result += 1
			flag = true
		} else {
			if !flag {
				continue
			} else {
				break
			}
		}
	}
	return result
}

// @leet end

func main() {}
