package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	assert.Equal(t, "olleh", reverseString("hello"))
	assert.Equal(t, "a", reverseString("a"))
	assert.Equal(t, "olleh", reverseString2("hello"))
	assert.Equal(t, "a", reverseString2("a"))
}
