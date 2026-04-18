/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


func removeNthFromEnd(head *ListNode, n int) *ListNode {
	forward := head
	for range n {
		forward = forward.Next
	}

	var prev *ListNode
	curr := head
	for ; forward != nil; forward = forward.Next {
		prev = curr
		curr = curr.Next
	}

	if curr == head {
		return head.Next
	}

	prev.Next = curr.Next
	return head
}
