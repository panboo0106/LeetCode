package main

// @leet start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{Next: nil}
	current := result
	add := 0
	for l1 != nil || l2 != nil {
		if l1 == nil {
			l1 = &ListNode{Next: nil}
		}
		if l2 == nil {
			l2 = &ListNode{Next: nil}
		}
		temp := current.Next
		current.Next = &ListNode{Val: (l1.Val + l2.Val + add) % 10}
		current.Next.Next = temp
		current = current.Next
		add = (l1.Val + l2.Val + add) / 10
		l1 = l1.Next
		l2 = l2.Next
	}
	if add == 1 {
		current.Next = &ListNode{Val: 1}
	}
	return result.Next
}

// @leet end

func main() {}
