package main

// @leet start
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	if n&(n-1) == 0 {
		return true
	} else {
		return false
	}
}

// @leet end

func main() {}
