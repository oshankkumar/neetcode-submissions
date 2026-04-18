/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    if head == nil {
		return nil
	}
	prev, next := head, head.Next
	for next != nil {
		temp := next.Next
		next.Next = prev
		prev = next
		next = temp
	}
    head.Next = nil
	return prev
}


