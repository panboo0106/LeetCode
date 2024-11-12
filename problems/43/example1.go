package main

import "strings"

// @leet start
// 大数相乘先把大数转化为切换存储
// 在通过竖式相乘算法
// 最后结果把大数转化为字符串
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	bnNum1 := NewBigNumber(num1)
	bnNum2 := NewBigNumber(num2)
	result := make([]int, len(bnNum1.digits)+len(bnNum2.digits))
	for i := 0; i < len(bnNum1.digits); i++ {
		for j := 0; j < len(bnNum2.digits); j++ {
			result[i+j] += bnNum1.digits[i] * bnNum2.digits[j]
		}
	}
	for i := 0; i < len(result)-1; i++ {
		result[i+1] += result[i] / 10
		result[i] %= 10
	}
	lastNonZero := len(result) - 1
	for lastNonZero >= 0 && result[lastNonZero] == 0 {
		lastNonZero--
	}
	var sb strings.Builder
	for i := lastNonZero; i >= 0; i-- {
		sb.WriteByte(byte(result[i]) + '0')
	}
	return sb.String()
}

type BigNumber struct {
	digits []int
}

func NewBigNumber(s string) *BigNumber {
	digits := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		digits[len(s)-1-i] = int(s[i] - '0')
	}
	return &BigNumber{digits: digits}
}

// @leet end

func main() {}
