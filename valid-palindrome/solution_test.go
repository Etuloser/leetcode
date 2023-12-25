package validpalindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	testCase := []struct {
		s        string
		expected bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
	}
	for _, tc := range testCase {
		got1 := solution1(tc.s)
		assert.Equal(t, tc.expected, got1)
		got2 := solution2(tc.s)
		assert.Equal(t, tc.expected, got2)

	}
}
