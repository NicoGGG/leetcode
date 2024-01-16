from typing import List


class Solution:
    def findMedianSortedArrays1(self, nums1: List[int], nums2: List[int]) -> float:
        """
        Brute force solution by merging and then sorting the two arrays
        """
        merged_array = []
        for n1 in nums1:
            merged_array.append(n1)
        for n2 in nums2:
            merged_array.append(n2)
        merged_array.sort()
        l = len(merged_array)
        median = (
            merged_array[l // 2]
            if l % 2 == 1
            else (merged_array[l // 2 - 1] + merged_array[l // 2]) / 2
        )
        return median

    def findMedianSortedArrays2(self, nums1: List[int], nums2: List[int]) -> float:
        """
        Time and memory optimal solution by taking advantage of the fact that both arrays are already sorted
        """
        merged_array = []
        while nums1 and nums2:
            if nums1[-1] >= nums2[-1]:
                merged_array.append(nums1.pop())
            else:
                merged_array.append(nums2.pop())
        while nums1:
            merged_array.append(nums1.pop())
        while nums2:
            merged_array.append(nums2.pop())
        # Two pointer method instead of time consuming arithmetic operation
        low = 0
        high = len(merged_array) - 1
        while low < high:
            low += 1
            high -= 1
        median = (
            merged_array[low]
            if low == high
            else (merged_array[low] + merged_array[high]) / 2
        )
        return median
