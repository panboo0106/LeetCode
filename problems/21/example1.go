package main

// @leet start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{Next: nil}
	current := result
	for {
		if list1 != nil && list2 != nil {
			if list1.Val <= list2.Val {
				current.Next = list1
				list1 = list1.Next
			} else {
				current.Next = list2
				list2 = list2.Next
			}
			current = current.Next
		} else if list1 != nil {
			current.Next = list1
			list1 = nil
		} else if list2 != nil {
			current.Next = list2
			list2 = nil
		} else {
			break
		}
	}
	return result.Next
}

// @leet end

func main() {}