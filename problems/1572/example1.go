package main

// @leet start
func diagonalSum(mat [][]int) int {
	result := 0
	length := len(mat)
	if length%2 != 0 {
		result -= mat[length/2][length/2]
	}
	for i := 0; i < length; i++ {
		result += mat[i][i] + mat[i][length-i-1]
	}
	return result
}

// @leet end

func main() {}
