/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import "iter"

func kthSmallest(root *TreeNode, k int) int {
	var val int
	for node := range BinaryTreeInOrderNodes(root) {
		if k--; k == 0 {
			val = node.Val
			break
		}
	}
	return val
}

func BinaryTreeInOrderNodes(root *TreeNode) iter.Seq[*TreeNode] {
	return func(yield func(*TreeNode) bool) {
		var stack []*TreeNode

		curr := root

		for len(stack) > 0 || curr != nil {
			for curr != nil {
				stack = append(stack, curr)
				curr = curr.Left
			}

			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if !yield(top) {
				break
			}
			curr = top.Right
		}
	}
}
