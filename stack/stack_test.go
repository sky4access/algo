package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {

	var stack Stack

	assert.True(t, stack.IsEmpty())
	stack.Push("1")
	assert.False(t, stack.IsEmpty())
	stack.Push("2")
	stack.Push("3")

	pop, ok := stack.Pop()
	assert.True(t, ok)
	assert.Equal(t, "3", pop)
	pop, ok = stack.Pop()
	assert.Equal(t, "2", pop)
	pop, ok = stack.Pop()
	assert.Equal(t, "1", pop)
	assert.True(t, stack.IsEmpty())
	pop, ok = stack.Pop()
	assert.False(t,  ok)
	assert.Equal(t, "", pop)

}