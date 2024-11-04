package main

import "sort"

// @leet start
func canMakeArithmeticProgression(arr []int) bool {
	if len(arr) <= 2 {
		return true
	}
	sort.Ints(arr)

	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != arr[1]-arr[0] {
			return false
		}
	}
	return true
}

// @leet end

func main() {}
