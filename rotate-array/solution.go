package rotatearray

func solution1(nums []int, k int) {
	// 最容易想到的方法是利用额外的内存空间分别append切片然后copy到原数组
	// 使用了append和copy，都是和时间线性相关的，时间复杂度为O(k+n)，k是常数，可以简化为O(n)
	// 使用了额外的内存空间来存放旋转后的数组，空间复杂度为O(n)
	temp := []int{}
	temp = append(temp, nums[len(nums)-k:]...)
	temp = append(temp, nums[:len(nums)-k]...)
	copy(nums, temp)
}

func solution2(nums []int, k int) {
	// 这种算法是上面算法的优化版本，但是不容易想到
	// 要点在于新数组的下标为什么是(i+k) mod n
	// 时间、空间复杂度同上
	newNums := make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}
	copy(nums, newNums)
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}
func solution3(nums []int, k int) {
	// 先翻转数组，再翻转k mod n
	// 时间复杂度为O(n)，空间复杂度为O(1)
	k %= len(nums)
	reverse(nums)
	reverse(nums[k:])
	reverse(nums[:k])
}

// TODO: solution4环状替代