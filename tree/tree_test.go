package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var root = &TreeNode{10,
&TreeNode{20,
&TreeNode{40,
&TreeNode{70, nil,nil},
nil},
&TreeNode{50,nil,nil}},
&TreeNode{30,nil,&TreeNode{60, nil,nil}}}

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

