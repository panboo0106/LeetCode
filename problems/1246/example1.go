package main

// @leet start
func addBinary(a string, b string) string {
	// 确保 a 字符串较长
	if len(b) > len(a) {
		a, b = b, a
	}

	// 将短字符串前面补0，使两个字符串等长
	for len(a) > len(b) {
		b = "0" + b
	}

	// 从低位开始逐位相加
	carry := 0
	var result []byte
	for i := len(a) - 1; i >= 0; i-- {
		sum := int(a[i]-'0') + int(b[i]-'0') + carry
		result = append(result, byte(sum%2+'0'))
		carry = sum / 2
	}

	// 如果最高位有进位，添加到结果中
	if carry == 1 {
		result = append(result, '1')
	}

	// 反转结果字符串
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}

// @leet end

func main() {}
