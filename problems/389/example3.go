func findTheDifference(s string, t string) byte {
	// 计算所有字符的ASCII值之和
	var sum1, sum2 byte
	for i := 0; i < len(s); i++ {
		sum1 += s[i]
	}
	for i := 0; i < len(t); i++ {
		sum2 += t[i]
	}
	// 差值就是被添加的字符
	return sum2 - sum1
}
