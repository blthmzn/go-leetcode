package largest_odd_number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestOddNumber(t *testing.T) {
	tt := []struct {
		name     string
		number   string
		expected string
	}{
		{
			name:     "case #1",
			number:   "52",
			expected: "5",
		},
		{
			name:     "case #2",
			number:   "4206",
			expected: "",
		},
		{
			name:     "case #2",
			number:   "35427",
			expected: "35427",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := largestOddNumber(tc.number)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
