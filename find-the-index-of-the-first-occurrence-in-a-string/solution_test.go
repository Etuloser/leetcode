package findtheindexofthefirstoccurrenceinastring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	taseCase := []struct {
		haystack string
		needle   string
		expected int
	}{
		{"sadbutsad", "sad", 0},
		{"leetcode", "leeto", -1},
	}
	for _, tc := range taseCase {
		got := solution1(tc.haystack, tc.needle)
		assert.Equal(t, tc.expected, got)
	}
}
