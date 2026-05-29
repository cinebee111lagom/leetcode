// Package linked_lists - LeetCode Problem 21: Merge Two Sorted Lists
// Merge two sorted linked lists and return as a sorted linked list.
// Time: O(n+m), Space: O(1)
package linked_lists

// MergeTwoLists merges two sorted linked lists.
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}

	return dummy.Next
}