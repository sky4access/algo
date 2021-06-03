package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMeetingRoom(t *testing.T) {
	result := minMeetingRooms([][]int{{0, 30}, {5, 10}, {15, 20}})
	assert.Equal(t, 2, result)

	result = minMeetingRooms([][]int{{7, 10}, {2, 4}})
	assert.Equal(t, 1, result)
}
