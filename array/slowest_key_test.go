package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlowestKey(t *testing.T) {
	result := slowestKey([]int{9, 29, 49, 50}, "cbcd")
	assert.Equal(t, uint8('c'), result)
}
