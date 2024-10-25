package main

// @leet start
func maxScore(s string) int {
	totalOnces := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			totalOnces++
		}
	}
	result := 0
	leftZeors := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '0' {
			leftZeors++
		} else {
			totalOnces--
		}
		if leftZeors+totalOnces > result {
			result = leftZeors + totalOnces
		}
	}
	return result
}

// @leet end

func main() {}
