package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFib(t *testing.T) {
	res := fibMaster(10)
	assert.Equal(t, int64(55), res)
	res = fibMaster(20)
	assert.Equal(t, int64(6765), res)
}
