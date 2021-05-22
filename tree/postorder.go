package tree

// left, right, root
func postOrder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	postOrder(root.Left, result)
	postOrder(root.Right, result)
	*result = append(*result, root.Val)
}

// left, right, root
func postorderTraversal(root *TreeNode) []int {
	var result []int
	postOrder(root, &result)
	return result
}