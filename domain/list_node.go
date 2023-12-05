package domain

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Reverse() (prev *ListNode) {
	curr := l

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return
}

func Equal(l1 *ListNode, l2 *ListNode) bool {
	if l1 == nil && l2 == nil {
		return true
	}

	if (l1 == nil || l2 == nil) || (l1.Val != l2.Val) {
		return false
	}

	return Equal(l1.Next, l2.Next)
}
