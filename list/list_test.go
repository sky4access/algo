package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseLis(t *testing.T) {

	result := reverseList(&ListNode{1,
		&ListNode{2, &ListNode{3,
			&ListNode{4, &ListNode{5, nil}}}}})
	assert.Equal(t, &ListNode{5,
		&ListNode{4, &ListNode{3,
			&ListNode{2, &ListNode{1, nil}}}}}, result)
}

func TestReverseBetween(t *testing.T) {
	result := reverseBetween(&ListNode{1,
		&ListNode{2, &ListNode{3,
			&ListNode{4, &ListNode{5, nil}}}}}, 2, 4)
	assert.Equal(t, &ListNode{1,
		&ListNode{4, &ListNode{3,
			&ListNode{2, &ListNode{5, nil}}}}}, result)
}

func TestMergeTwoLists(t *testing.T) {
	result := mergeTwoLists(&ListNode{1, &ListNode{2, &ListNode{4, nil}}},
		&ListNode{1, &ListNode{3, &ListNode{4, nil}}})
	assert.Equal(t, &ListNode{1,
		&ListNode{1,
			&ListNode{2,
				&ListNode{3,
					&ListNode{4,
						&ListNode{4, nil}}}}}}, result)
}
