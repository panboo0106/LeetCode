// 解法2：从后往前合并，不需要临时数组
func merge(nums1 []int, m int, nums2 []int, n int) {
	p1, p2 := m-1, n-1
	current := m + n - 1

	// 从后往前比较并填充
	for p2 >= 0 { // 只需要考虑 nums2 没有处理完的情况
		if p1 >= 0 && nums1[p1] > nums2[p2] {
			nums1[current] = nums1[p1]
			p1--
		} else {
			nums1[current] = nums2[p2]
			p2--
		}
		current--
	}
}

