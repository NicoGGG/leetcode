from typing import List


def twoSum1(numbers: List[int], target: int) -> List[int]:
    """
    Brute force. Way too slow but works
    """
    for i, n in enumerate(numbers):
        for j, m in enumerate(numbers[i + 1 :]):
            if n + m > target:
                break
            if n + m == target:
                return [i + 1, j + i + 2]


def twoSum2(numbers: List[int], target: int) -> List[int]:
    """
    Optimal solution. Similar to twoSum
    """
    m: dict[int, int] = {}
    for i, n in enumerate(numbers):
        c = m.get(target - n)
        if c:
            return [c, i + 1]
        else:
            m[n] = i + 1
    return []


nums1 = [2, 7, 11, 15]
target1 = 9
nums2 = [2, 3, 4]
target2 = 6
nums3 = [-1, 0]
target3 = -1

num4file = open("input.txt", "r")
from ast import literal_eval

nums4 = literal_eval(num4file.readline())
target4 = 2

print("Solution 1")
print([1, 2], twoSum1(nums1, target1))
print([1, 3], twoSum1(nums2, target2))
print([1, 2], twoSum1(nums3, target3))
# print([29999, 30000], twoSum1(nums4, target4))
print("\nSolution 2")
print([1, 2], twoSum2(nums1, target1))
print([1, 3], twoSum2(nums2, target2))
print([1, 2], twoSum2(nums3, target3))
print([29999, 30000], twoSum2(nums4, target4))
