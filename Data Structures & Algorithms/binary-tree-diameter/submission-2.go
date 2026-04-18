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

	getDepth(root, func(l int) {
		d = max(d, l)
	})

	return d
}

func getDepth(root *TreeNode, visit func(longestPath int)) int {
	if root == nil {
		return 0
	}

	ld := getDepth(root.Left, visit)
	rd := getDepth(root.Right, visit)
	visit(ld + rd)
	return 1 + max(ld, rd)
}
