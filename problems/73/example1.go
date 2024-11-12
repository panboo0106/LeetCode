package main

// @leet start
func setZeroes(matrix [][]int) {
	rows := make([]bool, len(matrix))
	columns := make([]bool, len(matrix[0]))
	for i, row := range matrix {
		for j := range row {
			if matrix[i][j] == 0 {
				rows[i] = true
				columns[j] = true
			}
		}
	}
	for i, row := range matrix {
		for j := range row {
			if rows[i] == true || columns[j] == true {
				matrix[i][j] = 0
			}
		}
	}
}

// @leet end

func main() {}
