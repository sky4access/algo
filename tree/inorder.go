package tree

// left, root, right
func inorderTraversal(root *TreeNode) []int {
	var result []int
	inOrder(root, &result)
	return result
}

func inOrder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	inOrder(root.Left, result)
	*result = append(*result, root.Val)
	inOrder(root.Right, result)
}
