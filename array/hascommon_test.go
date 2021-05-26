package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasCommon(t *testing.T) {
	assert.True(t, hasCommon([]int{1, 2, 3}, []int{3, 4, 5}))
	assert.False(t, hasCommon([]int{1, 2, 3}, []int{4, 5, 6}))
	assert.False(t, hasCommon([]int{1, 2, 3}, []int{}))
}
