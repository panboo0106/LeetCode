func hIndex(citations []int) int {
	n := len(citations)
	// 创建计数数组
	counts := make([]int, n+1)

	// 统计引用次数分布
	for _, c := range citations {
		if c >= n {
			counts[n]++
		} else {
			counts[c]++
		}
	}

	// 从高到低累计论文数量
	total := 0
	for i := n; i >= 0; i-- {
		total += counts[i]
		if total >= i {
			return i
		}
	}

	return 0
}
