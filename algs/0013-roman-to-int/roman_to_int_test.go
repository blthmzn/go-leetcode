package roman_to_int

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanToInt(t *testing.T) {
	tt := []struct {
		name     string
		roman    string
		expected int
	}{
		{
			name:     "case #1",
			roman:    "III",
			expected: 3,
		},
		{
			name:     "case #2",
			roman:    "LVIII",
			expected: 58,
		},
		{
			name:     "case #3",
			roman:    "MCMXCIV",
			expected: 1994,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := romanToInt(tc.roman)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
