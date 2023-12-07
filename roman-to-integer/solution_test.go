package romantointeger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		s        string
		expected int
	}{
		{"III", 3},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}
	for _, tc := range testCases {
		got := solution1(tc.s)
		assert.Equal(t, tc.expected, got)
	}
}
