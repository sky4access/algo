package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryTree(t *testing.T) {

	root := &TreeNode{2, nil, nil}

	root.Insert(3)
	assert.Equal(t, &TreeNode{2, nil, &TreeNode{3, nil, nil}}, root)

	node := root.lookup(3)
	assert.NotNil(t, node)
	assert.Equal(t, &TreeNode{3, nil, nil}, node)
	root.Insert(1)
	assert.Equal(t, &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}, root)

	// nothing happens because 3 already exist
	root.Insert(3)
	assert.Equal(t, &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}, root)

}
