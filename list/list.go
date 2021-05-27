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
