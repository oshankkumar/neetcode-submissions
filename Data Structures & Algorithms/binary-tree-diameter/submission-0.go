/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	var d int

	var maxDepth func(root *TreeNode) int

	maxDepth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		l, r := maxDepth(root.Left), maxDepth(root.Right)
		d = max(d, l+r)
		return 1 + max(l, r)
	}

	maxDepth(root)

	return d
}
