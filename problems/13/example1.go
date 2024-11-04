package main

// @leet start
func romanToInt(s string) int {
	result := 0
	mapArr := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	for i := 0; i < len(s)-1; i++ {
		if string(s[i]) == "I" && (string(s[i+1]) == "V" || string(s[i+1]) == "X") {
			result += -mapArr[string(s[i])]
		} else if string(s[i]) == "X" && (string(s[i+1]) == "L" || string(s[i+1]) == "C") {
			result += -mapArr[string(s[i])]
		} else if string(s[i]) == "C" && (string(s[i+1]) == "D" || string(s[i+1]) == "M") {
			result += -mapArr[string(s[i])]
		} else {
			result += mapArr[string(s[i])]
		}
	}
	return result + mapArr[string(s[len(s)-1])]
}

// @leet end

func main() {}
