func merge(nums1 []int, m int, nums2 []int, n int) {
	// 注意浅拷贝
	tempArr := make([]int, m)
	copy(tempArr, nums1[:m])
	point1, point2 := 0, 0
	for point1 < m || point2 < n {
		if point1 != m && point2 != n {
			if tempArr[point1] <= nums2[point2] {
				nums1[point1+point2] = tempArr[point1]
				point1++
			} else {
				nums1[point1+point2] = nums2[point2]
				point2++
			}
		} else if point1 != m {
			nums1[point1+point2] = tempArr[point1]
			point1++
		} else {
			nums1[point1+point2] = nums2[point2]
			point2++
		}
	}
}

