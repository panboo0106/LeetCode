package base

import (
	"reflect"
	"testing"
)

func TestRandomAccess(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	result := randomAccess(nums)

	// 检查结果是否在数组中
	if !find(nums, result) {
		t.Errorf("randomAccess() = %v, 不在数组 %v 中", result, nums)
	}
}

func TestInsert(t *testing.T) {
	testCases := []struct {
		name     string
		arr      []int
		num      int
		index    int
		expected []int
	}{
		{"插入中间", []int{1, 2, 3, 4}, 5, 2, []int{1, 2, 5, 3, 4}},
		{"插入开头", []int{1, 2, 3}, 0, 0, []int{0, 1, 2, 3}},
		{"插入结尾", []int{1, 2, 3}, 4, 3, []int{1, 2, 3, 4}},
		{"索引越界", []int{1, 2, 3}, 4, 5, []int{1, 2, 3}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := insert(tc.arr, tc.num, tc.index)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("insert() = %v, 期望 %v", result, tc.expected)
			}
		})
	}
}

func TestExtend(t *testing.T) {
	arr := []int{1, 2, 3}
	result := extend(arr)
	expected := []int{1, 2, 3, 0}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("extend() = %v, 期望 %v", result, expected)
	}
}

func TestDelete(t *testing.T) {
	testCases := []struct {
		name     string
		arr      []int
		index    int
		expected []int
	}{
		{"删除中间", []int{1, 2, 3, 4, 5}, 2, []int{1, 2, 4, 5}},
		{"删除开头", []int{1, 2, 3}, 0, []int{2, 3}},
		{"删除结尾", []int{1, 2, 3}, 2, []int{1, 2}},
		{"索引越界", []int{1, 2, 3}, 5, []int{1, 2, 3}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := delete(tc.arr, tc.index)
			if !reflect.DeepEqual(result[:len(tc.expected)], tc.expected) {
				t.Errorf("delete() = %v, 期望 %v", result[:len(tc.expected)], tc.expected)
			}
		})
	}
}

func TestFind(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	testCases := []struct {
		name     string
		target   int
		expected bool
	}{
		{"找到元素", 3, true},
		{"未找到元素", 6, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := find(arr, tc.target)
			if result != tc.expected {
				t.Errorf("find() = %v, 期望 %v", result, tc.expected)
			}
		})
	}
}
