import (
	"slices"
)

func isValidBST(root *TreeNode) bool {
	vals := TraverseInOrder(root, nil)
	return slices.IsSortedFunc(vals, func(a int, b int) int {
		if a <= b {
			return -1
		}
		return 1
	})
}

func TraverseInOrder(root *TreeNode, acc []int) []int {
	if root == nil {
		return acc
	}

	acc = TraverseInOrder(root.Left, acc)
	acc = append(acc, root.Val)
	acc = TraverseInOrder(root.Right, acc)
	return acc
}
