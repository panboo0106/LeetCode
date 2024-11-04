package main

// @leet start
func judgeCircle(moves string) bool {
	byteArr := []byte(moves)
	mapArr := make(map[string]int, len(moves))
	if len(byteArr)%2 != 0 {
		return false
	}
	for i := 0; i < len(byteArr); i++ {
		mapArr[string(byteArr[i])]++
	}
	if mapArr["R"] == mapArr["L"] && mapArr["U"] == mapArr["D"] {
		return true
	} else {
		return false
	}
}

// @leet end

func main() {}
