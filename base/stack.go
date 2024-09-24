package base

import (
	"errors"
	"fmt"
)

// 栈相关操作
func stack_array() {
	// 数组实现
	var stack []int
	// 入栈
	stack = append(stack, 1)
	// 出栈
	pop := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	// 获取栈顶元素
	peek := stack[len(stack)-1]
	fmt.Println(peek, pop)
}

// Stack 结构体表示栈
type Stack struct {
	top  *Node
	size int
}

// Node 结构体表示链表中的节点
type Node struct {
	next  *Node
	value int
}

// Peek 方法用于查看栈顶元素但不移除它
func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	return s.top.value, nil
}

// IsEmpty 方法用于检查栈是否为空
func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

// Size 方法返回栈的大小
func (s *Stack) Size() int {
	return s.size
}

// Push 入栈
// 存在栈顶指针, 先入后出
func (s *Stack) Push(value int) {
	s.top = &Node{s.top, value}
	s.size++
}

// Pop 出栈
func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	value := s.top.value
	s.top = s.top.next
	s.size--
	return value, nil
}
