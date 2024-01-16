class Solution:
    def twoSum(self, nums, target):
        h = {}
        for i, n in enumerate(nums):
            if target - n in h.keys():
                return [h[target - n], i]
            h[n] = i