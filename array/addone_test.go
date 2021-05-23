package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddOne(t *testing.T) {

	arr := []int{1,2,3}
	result := addOne(arr)
	assert.Equal(t, []int{1,2,4}, result)


	arr = []int{1,9,9}
	result = addOne(arr)
	assert.Equal(t, []int{2,0,0}, result)

	arr = []int{9,9,9}
	result = addOne(arr)
	assert.Equal(t, []int{1,0,0,0}, result)
}


func TestAddNumber(t *testing.T) {

	arr := []int{1,2,3}
	result := addNumber(arr, 4)
	assert.Equal(t, []int{1,2,7}, result)


	arr = []int{1,9,9}
	result = addNumber(arr, 3)
	assert.Equal(t, []int{2,0,2}, result)

	arr = []int{9,9,9}
	result = addNumber(arr, 9)
	assert.Equal(t, []int{1,0,0,8}, result)
}