package removeduplicatesfromsortedarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemove(t *testing.T) {
	nums := []int{1, 1, 2}
	got := RemoveDup(nums)
	assert.Equal(t, 2, got)
	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	got = RemoveDup(nums)
	assert.Equal(t, 5, got)
}
