package last_word_length

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLastWord(t *testing.T) {
	tt := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "case #1",
			input:    "Hello World",
			expected: 5,
		},
		{
			name:     "case #2",
			input:    "   fly me   to   the moon  ",
			expected: 4,
		},
		{
			name:     "case #3",
			input:    "a",
			expected: 1,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := lengthOfLastWord(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
