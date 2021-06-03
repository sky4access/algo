package list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	first := head
	second := head.Next
	for second != nil {
		temp := second.Next
		second.Next = first
		first = second
		second = temp
	}
	head.Next = nil
	head = first
	return head
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

/*
Input: head = [1,2,3,4,5], left = 2, right = 4
Output: [1,4,3,2,5]
*/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	result := &ListNode{}
	result.Next = head

	preHead := result
	cur := result

	for i := 0; i < left; i++ {
		preHead = cur
		cur = cur.Next
	}

	newTail := cur // point to the final node in the reverse sub-list

	// preHead -> cur -> cur.Next
	cur = cur.Next

	// we try to reverse the sub-list
	// preHead -> oldHead ... newTail -> cur -> nextNode
	//               |           |
	//
	for i := left; i < right; i++ {
		oldHead := preHead.Next
		preHead.Next = cur

		nextNode := cur.Next
		cur.Next = oldHead

		cur = nextNode
		newTail.Next = nextNode
	}

	return result.Next
}

// time: O(n + m), space: O(n + m)
func mergeTwoListsRecurr(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val < l2.Val {
		l1.Next = mergeTwoListsRecurr(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoListsRecurr(l1, l2.Next)
		return l2
	}
}

// Time: O(n+m), Space: O(1)
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	var preHead = &ListNode{-1, nil}
	prev := preHead

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}

	if l1 == nil {
		prev.Next = l2
	} else {
		prev.Next = l1
	}
	return preHead.Next
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//Edge case
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	//make sure l1 is always lower
	if l1.Val > l2.Val {
		l1, l2 = l2, l1
	}

	result := l1
	prevNode := l1
	for l2 != nil {
		for l1 != nil && l1.Val <= l2.Val {
			prevNode = l1
			l1 = l1.Next
		}

		if l1 == nil {
			prevNode.Next = l2
			break
		}

		//update l1's prev Node to new L2 node
		prevNode.Next = l2

		//update l2 Node, that needs to be iterated
		l2 = l2.Next

		//update the newly added l1 's element's next to l1's actual next
		prevNode.Next.Next = l1

		//Now update PrevNode
		prevNode = prevNode.Next

	}
	return result
}

func plusOne(head *ListNode) *ListNode {
	var dummyhead = &ListNode{}
	dummyhead.Next = head
	notNine := dummyhead
	for head != nil {
		if head.Val != 9 {
			notNine = head
		}
		head = head.Next
	}
	notNine.Val++
	notNine = notNine.Next
	for notNine != nil {
		notNine.Val = 0
		notNine = notNine.Next
	}

	if dummyhead.Val != 0 {
		return dummyhead
	} else {
		return dummyhead.Next
	}
}

// T: O(n1 + n2)
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reverseList(l1)
	l2 = reverseList(l2)

	var head *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		// get the current values
		x1, x2 := 0, 0
		if l1 != nil {
			x1 = l1.Val
		}
		if l2 != nil {
			x2 = l2.Val
		}

		// current sum and carry
		val := (carry + x1 + x2) % 10
		carry = (carry + x1 + x2) / 10

		// update the result: add to front
		curr := &ListNode{val, nil}
		curr.Next = head
		head = curr

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry != 0 {
		curr := &ListNode{carry, nil}
		curr.Next = head
		head = curr
	}
	return head
}
