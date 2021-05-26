package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var root = &TreeNode{10,
	&TreeNode{20,
		&TreeNode{40,
			&TreeNode{70, nil, nil},
			nil},
		&TreeNode{50, nil, nil}},
	&TreeNode{30, nil, &TreeNode{60, nil, nil}}}

/*
			10
		   /  \
		 20	   30
		/ \      \
	   40  50     60
	  /
	 70
*/

// left, root, right
func TestInOrderTraversal(t *testing.T) {
	result := inorderTraversal(root)
	t.Log(result)
	assert.Equal(t, []int{70, 40, 20, 50, 10, 30, 60}, result)

}

// root, left, right
func TestPreorderTraversal(t *testing.T) {
	result := preorderTraversal(root)
	t.Log(result)
	assert.Equal(t, []int{10, 20, 40, 70, 50, 30, 60}, result)
}

// left, right, root
func TestPostorderTraversal(t *testing.T) {
	result := postorderTraversal(root)
	t.Log(result)
	assert.Equal(t, []int{70, 40, 50, 20, 60, 30, 10}, result)
}

func TestFindTarget(t *testing.T) {

	root := &TreeNode{5,
		&TreeNode{3,
			&TreeNode{2, nil, nil},
			&TreeNode{4, nil, nil}},
		&TreeNode{6,
			nil,
			&TreeNode{7, nil, nil}}}

	result := findTarget(root, 9)
	assert.True(t, result)

	root = &TreeNode{2,
		&TreeNode{1, nil, nil},
		&TreeNode{3, nil, nil}}
	result = findTarget(root, 4)
	assert.True(t, result)
}
