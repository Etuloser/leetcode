package removeduplicatesfromsortedarrayii

func RemoveDupII(nums []int) int {
	// 小于两个元素直接返回即可
	n := len(nums)
	if n <= 2 {
		return n
	}
	// 快指针遍历数组
	// 慢指针原地写新元素
	// 从index=2开始
	// 当快指针指向元素不等于上上个（slow-2）元素，slow要写一次并且右移一位
	// 全检查完了slow就是符合题意的所有元素长度
	slow, fast := 2, 2
	for fast < n {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
