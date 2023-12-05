package rotatearray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testSulotion(t *testing.T, solution func([]int, int)) {
	testCase := []struct {
		nums []int
		k    int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3},
		{[]int{-1, -100, 3, 99}, 2},
	}
	for i, tc := range testCase {
		solution(tc.nums, tc.k)
		switch i {
		case 0:
			assert.Equal(t, []int{5, 6, 7, 1, 2, 3, 4}, tc.nums)
		case 1:
			assert.Equal(t, []int{3, 99, -1, -100}, tc.nums)
		}
	}
}

func TestSulotion(t *testing.T) {
	testSulotion(t, solution1)
	testSulotion(t, solution2)
	testSulotion(t, solution3)
}
