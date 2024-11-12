package main

// @leet start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	result := &ListNode{Next: nil}
	current := result
	for {
		if head != nil {
			temp := current.Next
			current.Next = head
			head = head.Next
			current.Next.Next = temp

		} else {
			break
		}
	}
	return result.Next
}

// @leet end

func main() {}
