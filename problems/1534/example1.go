package main

// @leet start

func countGoodTriplets(arr []int, a int, b int, c int) int {
	result := 0
	for i := 0; i < len(arr)-2; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			if abs(arr[i]-arr[j]) > a {
				continue
			}
			for k := j + 1; k < len(arr); k++ {
				if abs(arr[i]-arr[j]) <= a && abs(arr[j]-arr[k]) <= b && abs(arr[i]-arr[k]) <= c {
					result++
				}
			}
		}
	}
	return result
}

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

// @leet end

func main() {}
