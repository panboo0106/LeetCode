package main

// @leet start
func lemonadeChange(bills []int) bool {
	haspMap := make(map[int]int)
	for i := 0; i < len(bills); i++ {
		switch bills[i] {
		case 20:
			haspMap[20]++
			if haspMap[10] >= 1 && haspMap[5] >= 1 {
				haspMap[10]--
				haspMap[5]--
			} else if haspMap[5] >= 3 {
				haspMap[5] -= 3
			} else {
				return false
			}

		case 10:
			haspMap[10]++
			if haspMap[5] >= 1 {
				haspMap[5]--
			} else {
				return false
			}

		default:
			haspMap[5]++
		}
	}
	return true
}

// @leet end

func main() {}
