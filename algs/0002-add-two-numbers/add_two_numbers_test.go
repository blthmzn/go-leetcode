package add_two_numbers

import (
	"testing"

	"github.com/blthmzn/go-leetcode/domain"
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
			name: "case #1",
			l1: &domain.ListNode{
				Val: 2,
				Next: &domain.ListNode{
					Val: 4,
					Next: &domain.ListNode{
						Val: 3,
					},
				},
			},
			l2: &domain.ListNode{
				Val: 5,
				Next: &domain.ListNode{
					Val: 6,
					Next: &domain.ListNode{
						Val: 4,
					},
				},
			},
			expected: &domain.ListNode{
				Val: 7,
				Next: &domain.ListNode{
					Val: 0,
					Next: &domain.ListNode{
						Val: 8,
					},
				},
			},
		},
		{
			name: "case #2",
			l1: &domain.ListNode{
				Val: 0,
			},
			l2: &domain.ListNode{
				Val: 0,
			},
			expected: &domain.ListNode{
				Val: 0,
			},
		},
		{
			name: "case #3",
			l1: &domain.ListNode{
				Val: 9,
				Next: &domain.ListNode{
					Val: 9,
					Next: &domain.ListNode{
						Val: 9,
						Next: &domain.ListNode{
							Val: 9,
							Next: &domain.ListNode{
								Val: 9,
								Next: &domain.ListNode{
									Val: 9,
									Next: &domain.ListNode{
										Val: 9,
									},
								},
							},
						},
					},
				},
			},
			l2: &domain.ListNode{
				Val: 9,
				Next: &domain.ListNode{
					Val: 9,
					Next: &domain.ListNode{
						Val: 9,
						Next: &domain.ListNode{
							Val: 9,
						},
					},
				},
			},
			expected: &domain.ListNode{
				Val: 8,
				Next: &domain.ListNode{
					Val: 9,
					Next: &domain.ListNode{
						Val: 9,
						Next: &domain.ListNode{
							Val: 9,
							Next: &domain.ListNode{
								Val: 0,
								Next: &domain.ListNode{
									Val: 0,
									Next: &domain.ListNode{
										Val: 0,
										Next: &domain.ListNode{
											Val: 1,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := addTwoNumbers(tc.l1, tc.l2)
			assert.True(t, domain.Equal(tc.expected, res))
		})
	}
}
