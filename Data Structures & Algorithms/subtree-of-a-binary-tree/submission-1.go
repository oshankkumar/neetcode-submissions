/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func isEqualTree(p *TreeNode, q *TreeNode) bool {
	if p != nil && q != nil && p.Val == q.Val {
		return isEqualTree(p.Left, q.Left) && isEqualTree(p.Right, q.Right)
	}
	return p == nil && q == nil
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root != nil && subRoot != nil {
		if root.Val == subRoot.Val {
			return isEqualTree(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
		}
		return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
	}
	return root == nil && subRoot == nil
}