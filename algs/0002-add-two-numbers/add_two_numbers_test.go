package add_two_numbers

import (
	"testing"

	domain "github.com/blthmzn/go-leetcode/domain/list_domain"
	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	tt := []struct {
		name     string
		l1       *domain.ListNode
		l2       *domain.ListNode
		expected *domain.ListNode
	}{
		{
			name:     "case #1",
			l1:       domain.Create([]int{2, 4, 3}),
			l2:       domain.Create([]int{5, 6, 4}),
			expected: domain.Create([]int{7, 0, 8}),
		},
		{
			name:     "case #2",
			l1:       domain.Create([]int{0}),
			l2:       domain.Create([]int{0}),
			expected: domain.Create([]int{0}),
		},
		{
			name:     "case #3",
			l1:       domain.Create([]int{9, 9, 9, 9, 9, 9, 9}),
			l2:       domain.Create([]int{9, 9, 9, 9}),
			expected: domain.Create([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := addTwoNumbers(tc.l1, tc.l2)
			assert.True(t, domain.Equal(tc.expected, res))
		})
	}
}
