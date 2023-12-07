package single_number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleNumber(t *testing.T) {
	tt := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "case #1",
			nums:     []int{2, 2, 1},
			expected: 1,
		},
		{
			name:     "case #2",
			nums:     []int{4, 1, 2, 1, 2},
			expected: 4,
		},
		{
			name:     "case #3",
			nums:     []int{1},
			expected: 1,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := singleNumber(tc.nums)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
