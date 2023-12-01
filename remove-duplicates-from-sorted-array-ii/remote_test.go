package removeduplicatesfromsortedarrayii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDupII(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	got := RemoveDupII(nums)
	assert.Equal(t, 5, got)
	nums = []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	got = RemoveDupII(nums)
	assert.Equal(t, 7, got)
}
