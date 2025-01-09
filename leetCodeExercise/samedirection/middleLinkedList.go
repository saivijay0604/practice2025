package leetcode

import "leetcode/practice2025/config"

// head is the liked list
// time complexity - o(n)

func MiddleNode(head *config.ListNode) *config.ListNode {

	if head == nil {
		return nil
	}
	fast := head // fast will move 2 nodes at at a time
	slow := head // slow will move 1 nodes at a time

	for fast != nil && fast.Next != nil {

		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
