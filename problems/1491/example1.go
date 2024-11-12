package main

import "sort"

// @leet start
func average(salary []int) float64 {
	var sum float64
	sort.Ints(salary)
	for i := 1; i < len(salary)-1; i++ {
		sum += float64(salary[i])
	}
	result := sum / float64(len(salary)-2)
	return result
}

// @leet end

func main() {}
