package main

// @leet start
func maximumWealth(accounts [][]int) int {
	result := -1
	for i := 0; i < len(accounts); i++ {
		sum := 0
		for j := 0; j < len(accounts[i]); j++ {
			sum += accounts[i][j]
		}
		if sum > result {
			result = sum
		}
	}
	return result
}

// @leet end

func main() {}
