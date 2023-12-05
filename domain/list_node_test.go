package domain_test

import (
	"testing"

	"github.com/blthmzn/go-leetcode/domain"
	"github.com/stretchr/testify/assert"
)

func TestListDomain_Reverse(t *testing.T) {
	tt := []struct {
		name     string
		list     *domain.ListNode
		expected *domain.ListNode
	}{
		{
			name: "case #1",
			list: &domain.ListNode{
				Val: 1,
				Next: &domain.ListNode{
					Val: 2,
					Next: &domain.ListNode{
						Val: 3,
						Next: &domain.ListNode{
							Val: 4,
						},
					},
				},
			},
			expected: &domain.ListNode{
				Val: 4,
				Next: &domain.ListNode{
					Val: 3,
					Next: &domain.ListNode{
						Val: 2,
						Next: &domain.ListNode{
							Val: 1,
						},
					},
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tc.list = tc.list.Reverse()
			assert.True(t, domain.Equal(tc.list, tc.expected))
		})
	}
}
