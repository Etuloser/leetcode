package removeelement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 3
	got := RemoveElement(nums, val)
	assert.Equal(t, got, 2)
	nums1 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val1 := 2
	got1 := RemoveElement(nums1, val1)
	assert.Equal(t, got1, 5)
}
