package main

import "strings"

// @leet start
func intToRoman(num int) string {
	// 预定义罗马数字映射表，包括特殊情况和基本符号
	values := []struct {
		value  int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	// 存储结果的罗马数字字符串
	var result strings.Builder

	// 遍历映射表
	for _, item := range values {
		// 当前数字大于等于当前值时，重复添加对应符号
		for num >= item.value {
			result.WriteString(item.symbol)
			num -= item.value
		}
	}

	return result.String()
}

// @leet end

func main() {}
