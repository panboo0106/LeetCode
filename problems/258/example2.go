package main

// @leet start
func addDigits(num int) int {
	if num <= 9 {
		return num
	}
	result := 0
	for num > 0 {
		result += num % 10
		num = num / 10
	}
	num = addDigits(result)

	return num
}

// @leet end

func main() {}
