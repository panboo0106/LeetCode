package base

import (
	"errors"
)

// queue 结构queue
type queue struct {
	rear  *node
	front *node
	size  int
}

// Node 结构体表示链表中的节点
type node struct {
	next  *node
	value int
}

// Peek 方法用于查看栈顶元素但不移除它
func (s *queue) Front() (int, error) {
	if s.isEmpty() {
		return 0, errors.New("stack is empty")
	}
	return s.front.value, nil
}

// IsEmpty 方法用于检查栈是否为空
func (s *queue) isEmpty() bool {
	return s.size == 0
}

// Size 方法返回栈的大小
func (s *queue) Size() int {
	return s.size
}

// enqueue 入队
func (s *queue) enqueue(value int) {
	newNode := &node{value: value}
	if s.isEmpty() {
		s.rear = newNode
		s.front = newNode
	} else {
		s.rear.next = newNode
		s.rear = newNode
	}
	s.size++
}

// dequeue 出队
func (s *queue) dequeue() (int, error) {
	if s.isEmpty() {
		return 0, errors.New("queue is empty")
	}
	value := s.front.value
	s.front = s.front.next
	s.size--
	return value, nil
}

/////////// 循环数组实现 ///////////

var (
	array       []int
	rear, front int
)

func enqueue(value int) error {
	if (rear+1)%len(array) == front {
		return errors.New("queue is full")
	}
	array[rear] = value
	rear = (rear + 1) % len(array)
	return nil
}

func dequeue() interface{} {
	if rear == front {
		return errors.New("queue is empty")
	}
	value := array[front]
	front = (front + 1) % len(array)
	return value
}
