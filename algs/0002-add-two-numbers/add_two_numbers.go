package add_two_numbers

import "github.com/blthmzn/go-leetcode/domain"

// input: l1 = [2,4,3], l2 = [5,6,4]
// output: [7,0,8]
func addTwoNumbers(l1 *domain.ListNode, l2 *domain.ListNode) *domain.ListNode {
	res := &domain.ListNode{}
	curr := res
	var current int

	for l1 != nil || l2 != nil || current > 0 {
		sum := current

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		current = sum / 10

		curr.Next = &domain.ListNode{Val: sum % 10}
		curr = curr.Next
	}

	return res.Next
}
