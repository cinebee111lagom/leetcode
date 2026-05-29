// Package linked_lists - LeetCode Problem 83: Remove Duplicates from Sorted List
// Given the head of a sorted linked list, delete all duplicates so each element appears only once.
// Time: O(n), Space: O(1)
package linked_lists

// DeleteDuplicates removes duplicates from a sorted linked list.
func DeleteDuplicates(head *ListNode) *ListNode {
	current := head
	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head
}