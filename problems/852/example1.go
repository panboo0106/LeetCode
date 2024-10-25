package main

// @leet start
func peakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)-1

	for left < right {
		mid := (left + right) / 2
		if mid == 0 {
			if arr[0] > arr[1] {
				return 0
			}
			return 1
		}
		if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
			return mid
		} else if arr[mid] < arr[mid+1] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

// @leet end

func main() {}
