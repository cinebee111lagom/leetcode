// Package linked_lists - LeetCode Problem 19: Remove Nth Node From End
// Given the head of a linked list, remove the nth node from the end and return the head.
// Time: O(n), Space: O(1)
package linked_lists

// RemoveNthFromEnd removes the nth node from the end of the list.
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	fast := dummy
	slow := dummy

	for i := 0; i <= n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}