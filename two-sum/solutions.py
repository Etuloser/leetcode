from typing import List


class Solutions:
    def towSum(self, nums: List[int], target: int) -> List[int]:
        hashtable = {}
        for i in range(len(nums)):
            if target - nums[i] in hashtable:
                return [hashtable[target - nums[i]], i]
            else:
                hashtable[nums[i]] = i


# 哈希表，遍历数组，将当前k v记录到哈希表中
