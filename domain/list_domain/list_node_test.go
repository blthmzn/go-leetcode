package domain_test

import (
	"testing"

	domain "github.com/blthmzn/go-leetcode/domain/list_domain"
	"github.com/stretchr/testify/assert"
)

func TestListDomain_Reverse(t *testing.T) {
	tt := []struct {
		name     string
		list     *domain.ListNode
		expected *domain.ListNode
	}{
		{
			name:     "case #1",
			list:     domain.Create([]int{1, 2, 3, 4}),
			expected: domain.Create([]int{4, 3, 2, 1}),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tc.list = tc.list.Reverse()
			assert.True(t, domain.Equal(tc.list, tc.expected))
		})
	}
}

func TestListNode_Create(t *testing.T) {
	tt := []struct {
		name     string
		input    []int
		expected *domain.ListNode
	}{
		{
			name:  "case #1",
			input: []int{1, 2, 3},
			expected: &domain.ListNode{
				Val: 1,
				Next: &domain.ListNode{
					Val: 2,
					Next: &domain.ListNode{
						Val: 3,
					},
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			l := domain.Create(tc.input)
			assert.True(t, domain.Equal(l, tc.expected))
		})
	}
}
