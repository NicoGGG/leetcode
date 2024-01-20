from typing import List


def longestConsecutive1(nums: List[int]) -> int:
    """
    This is O(n*log(n)) so not optimal solution
    """
    s_nums = list(set(sorted(nums)))
    r = 0
    for i, n in enumerate(s_nums):
        if i == 0:
            # First round
            current_r = 1
            r = 1
            continue
        if s_nums[i - 1] == n - 1:
            current_r += 1
            if current_r >= r:
                r = current_r
        else:
            current_r = 1

    return r


def longestConsecutive2(nums: List[int]) -> int:
    """
    This is O(n) so optimal solution (with the help of neetcode video)
    """
    nums = set(nums)
    longest_sequence = 0
    for n in nums:
        tn = n
        if n - 1 not in nums:
            current_sequence = 1
            while tn + 1 in nums:
                current_sequence += 1
                tn += 1
            if current_sequence >= longest_sequence:
                longest_sequence = current_sequence

    return longest_sequence


a1 = [100, 4, 200, 1, 3, 2]
a2 = [0, 3, 7, 2, 5, 8, 4, 6, 0, 1]
a3 = [1, 2, 0, 1]
a4file = open("input.txt", "r")

from ast import literal_eval

a4 = literal_eval(a4file.read())
print("Solution 1")
print(4, longestConsecutive1(a1))
print(9, longestConsecutive1(a2))
print(3, longestConsecutive1(a3))

print("\nSolution 2")
print(4, longestConsecutive2(a1))
print(9, longestConsecutive2(a2))
print(3, longestConsecutive2(a3))
print(100000, longestConsecutive2(a4))
