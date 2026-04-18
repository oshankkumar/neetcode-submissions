/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	return NewTree(root).MaxDiameter()
}

func NewTree(root *TreeNode) *Tree {
	return &Tree{root: root}
}

type Tree struct {
	root        *TreeNode
	maxDiameter int
}

func (t *Tree) MaxDiameter() int {
	t.computeDepth(t.root)
	return t.maxDiameter
}

func (t *Tree) computeDepth(curr *TreeNode) int {
	if curr == nil {
		return 0
	}
	l := t.computeDepth(curr.Left)
	r := t.computeDepth(curr.Right)
	t.maxDiameter = max(t.maxDiameter, l+r)
	return 1 + max(l, r)
}
