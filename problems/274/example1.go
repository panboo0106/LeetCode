package main

import "sort"

// @leet start
func hIndex(citations []int) int {
	sort.Slice(citations, func(i, j int) bool { return citations[i] > citations[j] })
	for i := 0; i < len(citations); i++ {
		if i+1 > citations[i] {
			return i
		}
	}
	return len(citations)
}

// @leet end

func main() {}
