package base

import "errors"

type ArrayDeque struct {
	items    []interface{}
	front    int
	rear     int
	size     int
	capacity int
}

func NewArrayDeque(capacity int) *ArrayDeque {
	return &ArrayDeque{
		items:    make([]interface{}, capacity),
		capacity: capacity,
		front:    0,
		rear:     0,
		size:     0,
	}
}

func (d *ArrayDeque) PushFront(item interface{}) error {
	if d.IsFull() {
		return errors.New("deque is full")
	}
	d.front = (d.front - 1 + d.capacity) % d.capacity // 防止负数
	d.items[d.front] = item
	d.size++
	return nil
}

func (d *ArrayDeque) PushBack(item interface{}) error {
	if d.IsFull() {
		return errors.New("deque is full")
	}
	d.items[d.rear] = item
	d.rear = (d.rear + 1) % d.capacity
	d.size++
	return nil
}

func (d *ArrayDeque) PopFront() (interface{}, error) {
	if d.IsEmpty() {
		return nil, errors.New("deque is empty")
	}
	item := d.items[d.front]
	d.items[d.front] = nil // 垃圾回收
	d.front = (d.front + 1) % d.capacity
	d.size--
	return item, nil
}

func (d *ArrayDeque) PopBack() (interface{}, error) {
	if d.IsEmpty() {
		return nil, errors.New("deque is empty")
	}
	item := d.items[d.rear]
	d.items[d.rear] = nil
	d.rear = (d.rear - 1 + d.capacity) % d.capacity
	d.size--
	return item, nil
}

func (d *ArrayDeque) IsEmpty() bool {
	return d.size == 0
}

func (d *ArrayDeque) IsFull() bool {
	return d.size == d.capacity
}

type Node1 struct {
	value interface{}
	prev  *Node1
	Next  *Node1
}

type Deque struct {
	front *Node1
	rear  *Node1
	size  int
}

func NewDeque() *Deque {
	return &Deque{}
}

func (d *Deque) PushFront(item interface{}) {
	newNode := &Node1{value: item}
	if d.IsEmpty() {
		d.front = newNode
		d.rear = newNode
	} else {
		newNode.Next = d.front
		d.front.prev = newNode
		d.front = newNode
	}
	d.size++
}

func (d *Deque) PushBack(item interface{}) {
	newNode := &Node1{value: item}
	if d.IsEmpty() {
		d.front = newNode
		d.rear = newNode
	} else {
		newNode.prev = d.rear
		d.rear.Next = newNode
		d.rear = newNode
	}
	d.size++
}

func (d *Deque) PopFront() (interface{}, bool) {
	if d.IsEmpty() {
		return nil, false
	}
	value := d.front.value
	d.front = d.front.Next
	if d.front == nil {
		d.rear = nil
	} else {
		d.rear.Next = nil
	}
	d.size--
	return value, true
}

func (d *Deque) PopBack() (interface{}, bool) {
	if d.IsEmpty() {
		return nil, false
	}
	value := d.rear.value
	d.rear = d.rear.prev
	if d.rear == nil {
		d.front = nil
	} else {
		d.rear.Next = nil
	}
	d.size--
	return value, true
}

// PeekFront returns the front element without removing it
func (d *Deque) PeekFront() (interface{}, bool) {
	if d.IsEmpty() {
		return nil, false
	}
	return d.front.value, true
}

// PeekRear returns the rear element without removing it
func (d *Deque) PeekRear() (interface{}, bool) {
	if d.IsEmpty() {
		return nil, false
	}
	return d.rear.value, true
}

func (d *Deque) IsEmpty() bool {
	return d.size == 0
}

func (d *Deque) Size() int {
	return d.size
}
