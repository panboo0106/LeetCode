package sort

import (
	"reflect"
	"testing"
)

var tests = []struct {
	name     string
	input    []int
	expected []int
}{
	{
		name:     "Already sorted",
		input:    []int{1, 2, 3, 4, 5},
		expected: []int{1, 2, 3, 4, 5},
	},
	{
		name:     "Reverse sorted",
		input:    []int{5, 4, 3, 2, 1},
		expected: []int{1, 2, 3, 4, 5},
	},
	{
		name:     "Mixed order",
		input:    []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3},
		expected: []int{1, 1, 2, 3, 3, 4, 5, 5, 6, 9},
	},
	{
		name:     "With duplicates",
		input:    []int{3, 3, 3, 2, 2, 1},
		expected: []int{1, 2, 2, 3, 3, 3},
	},
	{
		name:     "Empty slice",
		input:    []int{},
		expected: []int{},
	},
	{
		name:     "Single element",
		input:    []int{1},
		expected: []int{1},
	},
	{
		name:     "Large scale data",
		input:    []int{100, 99, 98, 97, 96, 95, 94, 93, 92, 91, 90, 89, 88, 87, 86, 85, 84, 83, 82, 81, 80, 79, 78, 77, 76, 75, 74, 73, 72, 71, 70, 69, 68, 67, 66, 65, 64, 63, 62, 61, 60, 59, 58, 57, 56, 55, 54, 53, 52, 51, 50, 49, 48, 47, 46, 45, 44, 43, 42, 41, 40, 39, 38, 37, 36, 35, 34, 33, 32, 31, 30, 29, 28, 27, 26, 25, 24, 23, 22, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100},
	},
	{
		name:     "Negative numbers",
		input:    []int{-5, -2, -8, 0, -1, 3, -3},
		expected: []int{-8, -5, -3, -2, -1, 0, 3},
	},
	{
		name:     "All zeros",
		input:    []int{0, 0, 0, 0, 0},
		expected: []int{0, 0, 0, 0, 0},
	},
	{
		name:     "Mixed positive and negative",
		input:    []int{-1, 5, -3, 2, 0, -4, 3},
		expected: []int{-4, -3, -1, 0, 2, 3, 5},
	},
	{
		name:     "Random large numbers",
		input:    []int{1000000, 10, 999999, 100, 5, 999998, 1000, 50},
		expected: []int{5, 10, 50, 100, 1000, 999998, 999999, 1000000},
	},
	{
		name:     "Repeated elements",
		input:    []int{1, 1, 1, 2, 2, 3, 3, 3, 1, 2, 3},
		expected: []int{1, 1, 1, 1, 2, 2, 2, 3, 3, 3, 3},
	},
}

func TestSelectSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SelectSort(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("selectSort() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestBubbleSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BubbleSort(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("bubbleSort() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestInsertSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := InsertSort(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("insertSort() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestQuickSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := QuickSort(tt.input, 0, len(tt.input)-1)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("quickSort() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// 基准测试
func BenchmarkBubbleSort(b *testing.B) {
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				BubbleSort(tt.input)
			}
		})
	}
}

func BenchmarkSelectSort(b *testing.B) {
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SelectSort(tt.input)
			}
		})
	}
}

func BenchmarkInsertSort(b *testing.B) {
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				InsertSort(tt.input)
			}
		})
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				QuickSort(tt.input, 0, len(tt.input)-1)
			}
		})
	}
}
