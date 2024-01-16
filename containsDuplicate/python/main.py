from typing import List


def containsDuplicate(nums: List[int]) -> bool:
    return len(set(nums)) != len(nums)
