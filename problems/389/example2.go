func findTheDifference(s string, t string) byte {
	var result byte
	// 对所有字符进行异或运算
	for i := 0; i < len(s); i++ {
		result ^= s[i]
	}
	for i := 0; i < len(t); i++ {
		result ^= t[i]
	}
	return result
}
