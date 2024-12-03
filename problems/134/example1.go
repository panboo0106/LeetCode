package main

// 贪心思想： 一旦发现当前路径不可行，立即尝试下一个起点
// @leet start
func canCompleteCircuit(gas []int, cost []int) int {
	start := 0
	current_surplus := 0
	total_deficit := 0
	for i := 0; i < len(gas); i++ {
		current_surplus += gas[i] - cost[i]
		if current_surplus < 0 {
			total_deficit += current_surplus
			current_surplus = 0
			start = i + 1
		}
	}
	if current_surplus+total_deficit >= 0 {
		return start
	} else {
		return -1
	}
}

// @leet end

func main() {}
