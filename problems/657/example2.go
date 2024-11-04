package main

// @leet start
func judgeCircle(moves string) bool {
	x, y := 0, 0

	// 遍历每个移动指令
	for _, move := range moves {
		switch move {
		case 'U':
			y++ // 向上移动
		case 'D':
			y-- // 向下移动
		case 'L':
			x-- // 向左移动
		case 'R':
			x++ // 向右移动
		}
	}

	// 判断是否回到原点 (0,0)
	return x == 0 && y == 0
}

// @leet end

func main() {}
