// Package linked_lists - LeetCode Problem 92: Reverse Linked List II
// Given the head of a linked list and two integers left and right,
// reverse the nodes from position left to position right and return the head.
// Time: O(n), Space: O(1)
package linked_lists

// ReverseBetween reverses nodes from left to right position.
func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy

	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}

	current := prev.Next
	for i := 0; i < right-left; i++ {
		next := current.Next
		current.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}

	return dummy.Next
}