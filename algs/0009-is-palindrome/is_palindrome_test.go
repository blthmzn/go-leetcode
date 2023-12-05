package is_palindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	tt := []struct {
		name     string
		input    int
		expected bool
	}{
		{
			name:     "case #1",
			input:    121,
			expected: true,
		},
		{
			name:     "case #2",
			input:    -121,
			expected: false,
		},
		{
			name:     "case #3",
			input:    10,
			expected: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, isPalindrome(tc.input))
		})
	}
}
