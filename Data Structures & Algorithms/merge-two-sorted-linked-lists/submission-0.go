/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pushBack(head **ListNode, tail **ListNode, val int) {
	if *head == nil {
		*head = &ListNode{Val: val}
		*tail = *head
		return
	}

	(*tail).Next = &ListNode{Val: val}
	(*tail) = (*tail).Next
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head, tail *ListNode

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			pushBack(&head, &tail, list1.Val)
			list1 = list1.Next
		} else {
			pushBack(&head, &tail, list2.Val)
			list2 = list2.Next
		}
	}

	pending := list1
	if list2 != nil {
		pending = list2
	}

	for ; pending != nil; pending = pending.Next {
		pushBack(&head, &tail, pending.Val)
	}

	return head
}
