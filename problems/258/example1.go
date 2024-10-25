package main

// @leet start
func addDigits(num int) int {
	for num > 9 {
		result := 0
		for num > 0 {
			result += num % 10
			num = num / 10
		}
		num = result
	}
	return num
}

// @leet end

func main() {}
