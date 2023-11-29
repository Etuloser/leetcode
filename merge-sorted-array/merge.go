package mergesortedarray

import (
	"sort"
)

func Merge(nums1 []int, m int, nums2 []int, n int) []int {
	// 最容易想到的方式
	// 直接复制num2的所有元素到num1后使用快排
	// 快排的时间复杂度是 O(nlogn)
	// 所以这种解法的时间复杂度是O((m+n)log(m+n))
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
	return nums1
}

func Merge2(nums1 []int, m int, nums2 []int, n int) []int {
	// 双指针解法
    // 由于两个切片是有序的
	// 可以将两个切片视作队列
	// 每次取切片头比大小
	// 小的放临时切片
	sorted := make([]int, 0, m+n)
	p1, p2 := 0, 0
	for {
		if p1 == m {
			sorted = append(sorted, nums2[p2:]...)
			break
		}
		if p2 == n {
			sorted = append(sorted, nums1[p1:]...)
		}

		if nums1[p1] < nums2[p2] {
			sorted = append(sorted, nums1[p1])
			p1++
		} else {
			sorted = append(sorted, nums2[p2])
			p2++
		}
	}

	copy(nums1, sorted)
	return nums1
}
