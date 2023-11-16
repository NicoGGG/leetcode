# [Median of Two Sorted Array](https://leetcode.com/problems/median-of-two-sorted-arrays/)

## Problem

Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

The overall run time complexity should be O(log (m+n)).

Example 1:

Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.
Example 2:

Input: nums1 = [1,2], nums2 = [3,4]
Output: 2.50000
Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.

Constraints:

nums1.length == m
nums2.length == n
0 <= m <= 1000
0 <= n <= 1000
1 <= m + n <= 2000
-106 <= nums1[i], nums2[i] <= 106

## Solution

`python -m unittest` to execute the tests (unittest must be installed globally or create a virtualenv for the exercice)

### O((n+m)log(n+m))

Merge the two array, sort the resulting array, and then get the median value: array[l//2 + 1] if l is odd, (array[l//2 - 1] + array[l//2])/2 if l is even.

This is working but it is not O(log(n+m)) because we do not take advantage of the fact that both array are already sorted here, wasting time and memory redoing a sort after the merge.

### O(log(m+n))

Traverse both array simultaneously and pop the highest value. At the end of one array, traverse the other and pop everything. The merged_array is reversed here, but since we are looking for the median value, it doesn't matter

### Optimal

To avoid doing arithmetic operation that can be costly, the two pointer method approach is used in the findMedianSortedArrays2 method.
