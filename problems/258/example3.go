package main

// @leet start
// 数字根特性：
// 数字能被9整除那么数字根为9
// 数字不能被9整除那么数字根为余数
func addDigits(num int) int {
	if num == 0 {
		return 0
	}
	if num%9 == 0 {
		return 9
	}
	return num % 9
}

// @leet end

func main() {}
