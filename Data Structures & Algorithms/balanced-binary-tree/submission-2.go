/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func isBalanced(root *TreeNode) bool {
	result := true
	traverseHeight(root, func(lh, rh int) {
		if abs(lh-rh) > 1 {
			result = false
		}
	})
	return result
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func traverseHeight(root *TreeNode, visit func(lh, rh int)) int {
	if root == nil {
		return 0
	}

	lh := traverseHeight(root.Left, visit)
	rh := traverseHeight(root.Right, visit)
	visit(lh, rh)
	return 1 + max(lh, rh)
}