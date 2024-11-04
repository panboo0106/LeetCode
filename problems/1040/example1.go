package main

// @leet start
func isRobotBounded(instructions string) bool {
	x, y := 0, 0
	direction := "y"
	flag := 1
	bytes := []byte(instructions)
	for k := 0; k < 4; k++ {
		for i := 0; i < len(instructions); i++ {
			switch bytes[i] {
			case 'G':
				if direction == "x" {
					x += flag
				}
				if direction == "y" {
					y += flag
				}
			case 'L':
				if direction == "y" {
					direction = "x"
					flag = -flag
				} else {
					direction = "y"
				}
			case 'R':
				if direction == "x" {
					direction = "y"
					flag = -flag
				} else {
					direction = "x"
				}
			}
		}
		if x == 0 && y == 0 {
			return true
		}

	}
	return false
}

// @leet end

func main() {}
