package validpalindrome

import "strings"

func solution1(s string) bool {
	// 思路是先去符号转换成小写, 折半比较首尾字符串是否相等
	// 时间复杂度, 空间复杂度都是O(|s|), s是字符串的长度
	var sgood string
	for i := 0; i < len(s); i++ {
		if isalnum(s[i]) {
			sgood += string(s[i])
		}
	}

	n := len(sgood)
	sgood = strings.ToLower(sgood)
	for i := 0; i < n/2; i++ {
		if sgood[i] != sgood[n-1-i] {
			return false
		}
	}
	return true
}

func solution2(s string) bool {
	// O(1)算法, 在原字符串上比较, 先将字符串转换为小写, 双指针指向首尾
	// 折半比较, 遇到不是字符串或数字的字符则跳过
	s = strings.ToLower(s)
	left, right := 0, len(s)-1

	for left < right {
		for left < right && !isalnum(s[left]) {
			left++
		}
		for left < right && !isalnum(s[right]) {
			right--
		}
		if left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
	}
	return true
}
func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
