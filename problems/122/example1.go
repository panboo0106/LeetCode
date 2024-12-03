package main

// @leet start
func maxProfit(prices []int) int {
	// 1. 状态定义
	// dp[i][0]表示第i天不持有股票的最大利润
	// dp[i][1]表示第i天持有股票的最大利润
	// 2. 状态转移方程
	// dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
	// 前一天不持有，前一天持有今天卖出
	// dp[i][1] = max(dp[i-1][1],dp[i-1][0]-prices[i])
	// 前一天持有，前一天不持有但今天买入
	// 3.初始化
	// dp[0][0] = 0
	// dp[0][1] = -prices[0]
	n := len(prices)
	if n <= 1 {
		return 0
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	// 因为当前的最大利润只和前一天的值相关所以只需要保留昨天的值即可
	// 把dp二维数组可以替换为两个变量 空间复杂度降为O（1）
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @leet end

func main() {}
