package main

// @leet start
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	if n == 1 {
		return true
	}
	for (n % 2) == 0 {
		n = n / 2
	}
	if n%2 != 0 && n != 1 {
		return false
	} else {
		return true
	}
}

// @leet end

func main() {}
