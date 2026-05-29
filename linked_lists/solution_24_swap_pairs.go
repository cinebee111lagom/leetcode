// Package linked_lists - LeetCode Problem 24: Swap Nodes in Pairs
// Given a linked list, swap every two adjacent nodes and return the head.
// Time: O(n), Space: O(1)
package linked_lists

// SwapPairs swaps every two adjacent nodes.
func SwapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	current := dummy

	for current.Next != nil && current.Next.Next != nil {
		first := current.Next
		second := current.Next.Next
		first.Next = second.Next
		second.Next = first
		current.Next = second
		current = first
	}
	return dummy.Next
}