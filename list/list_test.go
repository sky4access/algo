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

	result = mergeTwoListsRecurr(&ListNode{1, &ListNode{2, &ListNode{4, nil}}},
		&ListNode{1, &ListNode{3, &ListNode{4, nil}}})
	assert.Equal(t, &ListNode{1,
		&ListNode{1,
			&ListNode{2,
				&ListNode{3,
					&ListNode{4,
						&ListNode{4, nil}}}}}}, result)

	result = mergeTwoLists2(&ListNode{1, &ListNode{2, &ListNode{4, nil}}},
		&ListNode{1, &ListNode{3, &ListNode{4, nil}}})
	assert.Equal(t, &ListNode{1,
		&ListNode{1,
			&ListNode{2,
				&ListNode{3,
					&ListNode{4,
						&ListNode{4, nil}}}}}}, result)

}

func TestPlusOne(t *testing.T) {
	result := plusOne(&ListNode{1, &ListNode{2, &ListNode{3, nil}}})
	assert.Equal(t, &ListNode{1, &ListNode{2, &ListNode{4, nil}}}, result)

	result = plusOne(&ListNode{1, &ListNode{9, &ListNode{9, nil}}})
	assert.Equal(t, &ListNode{2, &ListNode{0, &ListNode{0, nil}}}, result)
}

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{7, &ListNode{2, &ListNode{4, &ListNode{3, nil}}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	result := addTwoNumbers(l1, l2)
	assert.Equal(t, &ListNode{7, &ListNode{8, &ListNode{0, &ListNode{7, nil}}}}, result)
}

func TestReverseList(t *testing.T) {
	root := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	result := reverseList(root)
	assert.Equal(t, &ListNode{3, &ListNode{2, &ListNode{1, nil}}}, result)
}
