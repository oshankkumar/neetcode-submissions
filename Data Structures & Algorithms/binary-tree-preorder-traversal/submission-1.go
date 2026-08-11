import "iter"

func preorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	for val := range PreorderTraversal(root) {
		result = append(result, val)
	}
	return result
}

func PreorderTraversal(root *TreeNode) iter.Seq[int] {
	return func(yield func(int) bool) {
		preorderTraverse(root, yield)
	}
}

func preorderTraverse(root *TreeNode, yield func(int) bool) bool {
	if root == nil {
		return true
	}
	return yield(root.Val) &&
		preorderTraverse(root.Left, yield) &&
		preorderTraverse(root.Right, yield)
}
