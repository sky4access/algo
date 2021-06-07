package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxUnits(t *testing.T) {
	result := maximumUnits([][]int{{1, 3}, {2, 2}, {3, 1}}, 4)
	assert.Equal(t, 8, result)
}
