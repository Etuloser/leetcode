package mergesortedarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerger(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0}
	nums2 := []int{2, 5}
	got := Merge(nums1, 3, nums2, 2)
	assert.Equal(t, got, []int{1, 2, 2, 3, 5})
}
func TestMerger2(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0}
	nums2 := []int{2, 5}
	got := Merge2(nums1, 3, nums2, 2)
	assert.Equal(t, got, []int{1, 2, 2, 3, 5})
}
