package tree

func preOrder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	*result = append(*result, root.Val)
	preOrder(root.Left, result)
	preOrder(root.Right, result)
}

// root, left, right
func preorderTraversal(root *TreeNode) []int {
	var result []int
	preOrder(root, &result)
	return result
}
