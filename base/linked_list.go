package base

// 链表结构体
type ListNode struct {
	Val  int       // 节点值
	Next *ListNode // 指向下一个节点指针
}

// 插入链表
// 在链表n0中插入n1
func Insert(n0, n1 *ListNode) *ListNode {
	if n0 == nil {
		return n1
	}
	if n1 == nil {
		return n0
	}
	p := n0.Next
	n0.Next = n1
	n1.Next = p
	return n0
}

// 删除链表
// 删除链表指定节点
func Delete(n0 *ListNode, index int) *ListNode {
	if n0 == nil {
		return nil
	}
	if index == 0 {
		return n0.Next
	}
	head := n0
	for i := 0; i < index-1 && n0.Next != nil; i++ {
		n0 = n0.Next
	}
	if n0.Next != nil {
		n0.Next = n0.Next.Next
	}
	return head
}

// 查找链表
// 查找指定节点,返回索引
func Find(n0 *ListNode, target int) int {
	index := 0
	for n0 != nil {
		if n0.Val == target {
			return index
		}
		n0 = n0.Next
		index++
	}

	return -1
}
