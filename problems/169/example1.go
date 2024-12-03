package main

// @leet start
func majorityElement(nums []int) int {
	result := 0
	max := 0
	mapArr := make(map[int]int)
	for _, v := range nums {
		mapArr[v]++
	}
	for k, v := range mapArr {
		if v > max {
			result = k
			max = v
		}
	}
	return result
}

// @leet end

func main() {}
