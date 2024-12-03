package main

// @leet start
func romanToInt(s string) int {
	bytes := []byte(s)
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	sum := 0
	for i := 0; i < len(bytes)-1; i++ {
		if bytes[i] == 'I' && (bytes[i+1] == 'V' || bytes[i+1] == 'X') {
			sum += -romanMap[bytes[i]]
		} else if bytes[i] == 'X' && (bytes[i+1] == 'L' || bytes[i+1] == 'C') {
			sum += -romanMap[bytes[i]]
		} else if bytes[i] == 'C' && (bytes[i+1] == 'D' || bytes[i+1] == 'M') {
			sum += -romanMap[bytes[i]]
		} else {
			sum += romanMap[bytes[i]]
		}
	}
	return sum + romanMap[bytes[len(bytes)-1]]
}

// @leet end

func main() {}
