from typing import List


class Solution:
    def move_zeros(self, nums: List[int]):
        n = len(nums)
        fast = slow = 0
        while fast < n:
            if nums[fast] != 0:
                nums[slow], nums[fast] = nums[fast], nums[slow]
                slow += 1
            fast += 1
# 解法思路 - 双指针
# 快指针在遇到不为0的数，和慢指针交换值，然后慢指针进1，快指针进1