package tree

// insert inserts num into t, assuming t is binary search tree
func (t *TreeNode) Insert(num int) {
	curr := t
	var prev *TreeNode
	for curr != nil && curr.Val != num {
		if curr.Val == num {
			break
		}
		prev = curr
		if num < curr.Val {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}
	curr = &TreeNode{num, nil, nil}
	if num < prev.Val {
		prev.Left = curr
	} else if num > prev.Val {
		prev.Right = curr
	}

}

// lookup returns TreeNode that has value of num
func (t *TreeNode) lookup(num int) *TreeNode {
	curr := t
	for curr != nil {
		if curr.Val == num {
			return curr
		}
		if num < curr.Val {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}
	return curr
}

// remove removes num from t, assuming t is binary search tree
func (t *TreeNode) remove(num int) {

}
