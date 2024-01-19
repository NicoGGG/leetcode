from typing import List


def search(nums: List[int], target: int) -> int:
    l = len(nums) - 1
    i = 0
    while i <= l:
        t = (i + l) // 2
        if nums[t] == target:
            return t
        if nums[t] < target:
            i = t + 1
        else:
            l = t - 1

    return -1


nums1 = [-1, 0, 3, 5, 9, 12]
target1 = 9
nums2 = [-1, 0, 3, 5, 9, 12]
target2 = 2
nums3 = [-1, 0, 3, 5, 9, 12]
target3 = 12
num4 = [1, 3]
target4 = 3

print(4, search(nums1, target1))
print(-1, search(nums2, target2))
print(5, search(nums3, target3))
print(1, search(num4, target4))
