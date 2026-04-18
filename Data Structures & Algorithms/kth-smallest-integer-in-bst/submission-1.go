/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func kthSmallest(root *TreeNode, k int) int {
	stack := list.New()
	curr := root
	var prev *TreeNode

	for k > 0 {
		for curr != nil {
			stack.PushBack(curr)
			curr = curr.Left
		}

		prev = stack.Remove(stack.Back()).(*TreeNode)
		curr = prev.Right
		k--
	}

	return prev.Val
}


