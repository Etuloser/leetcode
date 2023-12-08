package lengthoflastword

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		s        string
		expected int
	}{
		{"Hello World", 5},
		{"   fly me   to   the moon  ", 4},
		{"luffy is still joyboy", 6},
	}
	for _, tc := range testCases {
		ans := solution1(tc.s)
		assert.Equal(t, tc.expected, ans)
	}
}
