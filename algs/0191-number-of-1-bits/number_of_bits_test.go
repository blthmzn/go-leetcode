package number_of_bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHammingWeight(t *testing.T) {
	tt := []struct {
		name     string
		input    uint32
		expected int
	}{
		{
			name:     "case #1",
			input:    0b00000000000000000000000000001011,
			expected: 3,
		},
		{
			name:     "case #2",
			input:    0b00000000000000000000000010000000,
			expected: 1,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := hammingWeight(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
