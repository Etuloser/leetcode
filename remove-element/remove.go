package removeelement

func RemoveElement(nums []int, val int) int {
	// 双指针解法
	// left,right指针初始指向第一个元素
	// 右指针会遍历完数组，当右指针指向的值不为val时
	// 原地往nums[left]写v，然后左指针往右移一位
	// 也就是说[0, left]个元素都不等于val
	// 共计写了left个元素
	left := 0
	for _, v := range nums { // v即nums[right]
		if v != val {
			nums[left] = v
			left++
		}
	}
	return left
}

// TODO: 双指针优化解法