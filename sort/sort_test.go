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

// 基准测试
func BenchmarkBubbleSort(b *testing.B) {
	input := []int{5, 4, 3, 2, 1}
	for i := 0; i < b.N; i++ {
		BubbleSort(input)
	}
}

func BenchmarkSelectSort(b *testing.B) {
	input := []int{5, 4, 3, 2, 1}
	for i := 0; i < b.N; i++ {
		SelectSort(input)
	}
}

func BenchmarkInsertSort(b *testing.B) {
	input := []int{5, 4, 3, 2, 1}
	for i := 0; i < b.N; i++ {
		InsertSort(input)
	}
}
