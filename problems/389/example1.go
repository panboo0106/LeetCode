func findTheDifference(s string, t string) byte {
	mapStringS := make(map[rune]int, len(s))
	mapStringT := make(map[rune]int, len(t))
	// 对字符串使用range时，v是rune类型。
	// 因为Go的字符串是以UTF-8编码存储的，
	// range会自动解码UTF-8
	for _, v := range s {
		mapStringS[v]++
	}
	for _, v := range t {
		mapStringT[v]++
	}
	for _, v := range t {
		if _, ok := mapStringS[v]; !ok {
			return byte(v)
		}
		if mapStringS[v] != mapStringT[v] {
			return byte(v)
		}
	}
	return 0
}

