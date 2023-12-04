package majorityelement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
	}
	for _, tc := range testCases {
		result1 := solution(tc.nums)
		assert.Equal(t, tc.expected, result1)

		result2 := solution1(tc.nums)
		assert.Equal(t, tc.expected, result2)

		result3 := solution2(tc.nums)
		assert.Equal(t, tc.expected, result3)
	}
}
