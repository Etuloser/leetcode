package removeduplicatesfromsortedarray

func RemoveDup(nums []int) int {
	// nums长度为0直接返回0
	n := len(nums)
	if n == 0 {
		return 0
	}
	// 只有一个元素的时候不用比较
	// 所以定义的双指针从1开始
	slow := 1
	// 再定义另一个指针来遍历nums，与前一个元素比较
	// 不相等就nums[slow]原地写，然后slow++移动
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
