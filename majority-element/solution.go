package majorityelement

import "sort"

func solution(nums []int) int {
	// 最容易想到的是排序后返回n/2的数，一定是众数
	// 因为使用了sort，所以空间复杂度是O(logn)，时间复杂度为O(nlogn)
	n := len(nums)
	sort.Ints(nums)
	return nums[n/2]
}

func solution1(nums []int) int {
	// 定义一个哈希表来做元素计数，同时遍历数组时使用打擂台的形式来维护最大值，省去最后遍历哈希表
	// 遍历了nums，时间复杂度为O(n)，使用了哈希表，最多存n个元素，空间复杂度为O(n)
	hashtable := make(map[int]int)
	majority, cnt := 0, 0
	for _, v := range nums {
		hashtable[v]++
		if hashtable[v] > cnt {
			majority = v
			cnt = hashtable[v]
		}
	}
	return majority
}

func solution2(nums []int) int {
	// 摩尔投票法：解决在任意多候选人中（选票无序），选出得票最多的人
	// 1. 对抗阶段：分属两个候选人的选票两两抵消
	// 2. 计算阶段：计算对抗结果中留下的候选人选票是否有效
	major := 0
	count := 0
	for _, v := range nums {
		// count为0，即当前无当选，本次循环的元素当选
		if count == 0 {
			major = v
		}
		// 已有当选，选票相同+1，否则-1
		if major == v {
			count++
		} else {
			count--
		}
	}
	return major
}
