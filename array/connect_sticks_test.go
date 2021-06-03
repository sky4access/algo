package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectSticks(t *testing.T) {
	result := connectSticks([]int{2, 4, 3})
	assert.Equal(t, 14, result)
	result = connectSticks2([]int{1, 8, 3, 5})
	assert.Equal(t, 30, result)
}
