package main

// @leet start
func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) == 2 {
		return true
	}
	// 使用交叉相乘取代来避免除法（避免处理除数为0的情况）
	for i := 2; i < len(coordinates); i++ {
		if (coordinates[i-1][1]-coordinates[i-2][1])*(coordinates[i][0]-coordinates[i-1][0]) !=
			(coordinates[i][1]-coordinates[i-1][1])*(coordinates[i-1][0]-coordinates[i-2][0]) {
			return false
		}
	}
	return true
}

// @leet end

func main() {}
