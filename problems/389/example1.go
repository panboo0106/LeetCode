func findTheDifference(s string, t string) byte {
	mapStringS := make(map[rune]int, len(s))
	mapStringT := make(map[rune]int, len(t))
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

